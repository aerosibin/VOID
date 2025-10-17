package services

import (
	"VOID/config"
	"errors"
	"math"

	"VOID/internal/models"
	"gorm.io/gorm"
)

func AssignNearestDriver(db *gorm.DB, orderID uint) (*models.Vehicle, error) {
	var o models.Order
	if err := db.First(&o, orderID).Error; err != nil {
		return nil, err
	}
	var vehicles []models.Vehicle
	if err := db.Where("status = ?", "available").Find(&vehicles).Error; err != nil {
		return nil, err
	}
	if len(vehicles) == 0 {
		return nil, errors.New("no available vehicles")
	}
	var best *models.Vehicle
	bestDist := math.MaxFloat64
	for i := range vehicles {
		v := &vehicles[i]
		d := distanceKm(v.Latitude, v.Longitude, o.PickupLat, o.PickupLon)
		if d < bestDist {
			bestDist = d
			best = v
		}
	}
	if best == nil {
		return nil, errors.New("no driver")
	}
	best.Status = "busy"
	if err := db.Save(best).Error; err != nil {
		return nil, err
	}
	o.AssignedToID = &best.ID
	o.Status = config.OrderStatusAssigned
	o.DistanceKm = bestDist
	if err := db.Save(&o).Error; err != nil {
		return nil, err
	}
	return best, nil
}

func distanceKm(lat1, lon1, lat2, lon2 float64) float64 {
	dx := lat1 - lat2
	dy := lon1 - lon2
	return math.Sqrt(dx*dx+dy*dy) * 111
}
