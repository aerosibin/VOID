package api

import (
	"net/http"
	"strconv"

	"VOID/internal/cache"
	"VOID/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AssignFleet(c *gin.Context, db *gorm.DB, cache *cache.Cache) {
	idStr := c.Param("id")
	if idStr == "" {
		idStr = c.Query("order_id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}
	vehicle, err := services.AssignNearestDriver(db, uint(id))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"vehicle": vehicle})
}
