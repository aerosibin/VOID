package services

import (
	"errors"
	"time"

	"VOID/config"
	"VOID/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("default-secret")

func RegisterUser(db *gorm.DB, u *models.User) error {
	if u.Email == "" || u.Password == "" {
		return errors.New("email and password required")
	}
	if err := u.SetPassword(u.Password); err != nil {
		return err
	}
	return db.Create(u).Error
}

func LoginUser(db *gorm.DB, email, password string, cfg *config.Config) (string, error) {
	var u models.User
	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
		return "", err
	}
	if !u.CheckPassword(password) {
		return "", errors.New("invalid credentials")
	}
	if cfg != nil && cfg.JWT.Secret != "" {
		jwtKey = []byte(cfg.JWT.Secret)
	}
	claims := jwt.MapClaims{
		"email": u.Email,
		"uid":   u.ID,
		"role":  u.Role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtKey)
	return signed, err
}

func ValidateToken(tokenStr string, cfg *config.Config) (bool, error) {
	if cfg != nil && cfg.JWT.Secret != "" {
		jwtKey = []byte(cfg.JWT.Secret)
	}
	_, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
