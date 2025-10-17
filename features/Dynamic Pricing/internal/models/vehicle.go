package models

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	DriverID  uint
	Plate     string
	Status    string // "available", "busy"
	Latitude  float64
	Longitude float64
}
