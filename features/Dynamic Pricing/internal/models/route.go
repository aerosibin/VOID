package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	OrderID    uint
	Path       string // simple representation (JSON/WKT) for demo
	TravelCost float64
}
