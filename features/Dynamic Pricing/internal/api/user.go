package api

import (
	"VOID/config"
	"net/http"

	"VOID/internal/models"
	"VOID/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUser registers a new user (customer or driver). Expects JSON {name,email,password,role,latitude,longitude}
func RegisterUser(c *gin.Context, db *gorm.DB) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	if err := services.RegisterUser(db, &u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

// LoginUser: expects {email,password}, returns token
func LoginUser(c *gin.Context, db *gorm.DB, cfg *config.Config) {
	var req models.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	token, err := services.LoginUser(db, req.Email, req.Password, cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
