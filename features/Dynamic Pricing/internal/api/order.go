package api

import (
	"net/http"
	"strconv"

	"VOID/config"
	"VOID/internal/cache"
	"VOID/internal/models"
	"VOID/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context, db *gorm.DB, cache *cache.Cache, cfg *config.Config) {
	var in struct {
		UserID     uint    `json:"user_id"`
		PickupLat  float64 `json:"pickup_lat"`
		PickupLon  float64 `json:"pickup_lon"`
		DropoffLat float64 `json:"dropoff_lat"`
		DropoffLon float64 `json:"dropoff_lon"`
		Priority   string  `json:"priority"`
	}
	if err := c.BindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	order := &models.Order{
		UserID:     in.UserID,
		PickupLat:  in.PickupLat,
		PickupLon:  in.PickupLon,
		DropoffLat: in.DropoffLat,
		DropoffLon: in.DropoffLon,
		Priority:   in.Priority,
		Status:     config.OrderStatusPending,
	}
	if err := services.CreateOrder(db, order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create order"})
		return
	}
	price, breakdown := services.CalculateDynamicPrice(db, order.ID, cache, cfg)
	order.ComputedPrice = price
	_ = db.Save(order)

	c.JSON(http.StatusCreated, gin.H{"order": order, "price": price, "breakdown": breakdown})
}

func GetOrderStatus(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	status, err := services.GetOrderStatus(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}
