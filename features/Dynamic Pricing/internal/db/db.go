package db

import (
	"VOID/config"
	"VOID/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB connects to Postgres if cfg.Database.URL set, otherwise uses sqlite file "void.db"
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg != nil && cfg.Database.URL != "" {
		dsn := cfg.Database.URL
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	return gorm.Open(sqlite.Open("void.db"), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Order{},
		&models.Route{},
		&models.PricingHistory{},
	)
}
