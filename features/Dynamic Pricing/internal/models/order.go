package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint
	PickupLat     float64
	PickupLon     float64
	DropoffLat    float64
	DropoffLon    float64
	Priority      string // "normal" or "express"
	Status        string
	AssignedToID  *uint
	DistanceKm    float64
	ComputedPrice float64
}
