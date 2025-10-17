package api

import (
	"net/http"
	"strconv"

	"VOID/config"
	"VOID/internal/cache"
	"VOID/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CalculateDynamicPrice(c *gin.Context, db *gorm.DB, cache *cache.Cache, cfg *config.Config) {
	idStr := c.Param("order_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}
	price, breakdown := services.CalculateDynamicPrice(db, uint(id), cache, cfg)
	c.JSON(http.StatusOK, gin.H{"final_price": price, "breakdown": breakdown})
}
