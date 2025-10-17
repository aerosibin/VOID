package services

import (
	"fmt"
	"math"

	"VOID/config"
	"VOID/internal/cache"
	"VOID/internal/models"
	"VOID/internal/pricing"
	"gorm.io/gorm"
)

func CalculateDynamicPrice(db *gorm.DB, orderID uint, cacheClient *cache.Cache, cfg *config.Config) (float64, map[string]float64) {
	var o models.Order
	if err := db.First(&o, orderID).Error; err != nil {
		return 0, map[string]float64{"error": 1}
	}
	dist := o.DistanceKm
	if dist == 0 {
		dist = approxDistanceKm(o.PickupLat, o.PickupLon, o.DropoffLat, o.DropoffLon)
	}
	trafficScore := 0.3
	if cacheClient != nil && cacheClient.RedisClient != nil {
		if v, err := cacheClient.Get(fmt.Sprintf("traffic:%d", orderID)); err == nil && v != "" {
			var parsed float64
			fmt.Sscanf(v, "%f", &parsed)
			trafficScore = parsed
		}
	} else {
		facts := pricing.LoadFactors()
		trafficScore = facts.TrafficDensity
	}
	demandIndex := computeDemandIndex(db)
	pcfg := cfg.Pricing
	final := pricing.CalculatePrice(pcfg.BasePrice, dist, trafficScore, demandIndex, pcfg)
	ph := models.PricingHistory{
		OrderID:      orderID,
		BasePrice:    pcfg.BasePrice,
		DistanceKm:   dist,
		TrafficScore: trafficScore,
		DemandIndex:  demandIndex,
		FinalPrice:   final,
	}
	_ = db.Create(&ph).Error
	return math.Round(final*100) / 100, map[string]float64{
		"base":     pcfg.BasePrice,
		"distance": dist,
		"traffic":  trafficScore,
		"demand":   demandIndex,
		"final":    final,
	}
}

func approxDistanceKm(lat1, lon1, lat2, lon2 float64) float64 {
	dx := lat1 - lat2
	dy := lon1 - lon2
	return math.Sqrt(dx*dx+dy*dy) * 111
}

func computeDemandIndex(db *gorm.DB) float64 {
	var pending int64
	var avail int64
	_ = db.Model(&models.Order{}).Where("status = ?", config.OrderStatusPending).Count(&pending).Error
	_ = db.Model(&models.Vehicle{}).Where("status = ?", "available").Count(&avail).Error
	return float64(pending) / (float64(avail) + 1.0)
}
