package scripts

import (
	"fmt"

	"VOID/internal/models"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {
	u1 := models.User{Name: "Alice Driver", Email: "alice@drv", Role: "driver", Latitude: 12.97, Longitude: 77.59}
	_ = db.Create(&u1).Error
	v1 := models.Vehicle{DriverID: u1.ID, Plate: "KA01", Status: "available", Latitude: 12.97, Longitude: 77.59}
	_ = db.Create(&v1).Error

	u2 := models.User{Name: "Bob Driver", Email: "bob@drv", Role: "driver", Latitude: 12.98, Longitude: 77.58}
	_ = db.Create(&u2).Error
	v2 := models.Vehicle{DriverID: u2.ID, Plate: "KA02", Status: "available", Latitude: 12.98, Longitude: 77.58}
	_ = db.Create(&v2).Error

	fmt.Println("seeded drivers and vehicles")
}
