package models

import "time"

type PricingHistory struct {
	ID           uint `gorm:"primaryKey"`
	OrderID      uint
	BasePrice    float64
	DistanceKm   float64
	TrafficScore float64
	DemandIndex  float64
	FinalPrice   float64
	CreatedAt    time.Time
}
