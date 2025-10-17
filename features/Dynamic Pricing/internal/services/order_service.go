package services

import (
	"errors"

	"VOID/internal/models"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB, o *models.Order) error {
	if o.UserID == 0 {
		return errors.New("user required")
	}
	return db.Create(o).Error
}

func GetOrderStatus(db *gorm.DB, id uint) (string, error) {
	var o models.Order
	if err := db.First(&o, id).Error; err != nil {
		return "", err
	}
	return o.Status, nil
}
