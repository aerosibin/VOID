package api

import (
	"VOID/config"
	"VOID/internal/cache"
	"VOID/internal/ws"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cacheClient *cache.Cache, manager *ws.Manager, cfg *config.Config) {
	v1 := r.Group("/api/v1")

	v1.POST("/register", func(c *gin.Context) { RegisterUser(c, db) })
	v1.POST("/login", func(c *gin.Context) { LoginUser(c, db, cfg) })

	v1.POST("/orders", func(c *gin.Context) { CreateOrder(c, db, cacheClient, cfg) })
	v1.GET("/orders/:id/status", func(c *gin.Context) { GetOrderStatus(c, db) })
	v1.POST("/orders/:id/assign", func(c *gin.Context) { AssignFleet(c, db, cacheClient) })

	v1.GET("/pricing/:order_id", func(c *gin.Context) { CalculateDynamicPrice(c, db, cacheClient, cfg) })

	v1.GET("/route", func(c *gin.Context) { ComputeRoute(c) })

	v1.GET("/ws", func(c *gin.Context) { WebSocketHandler(manager)(c.Writer, c.Request) })
}
