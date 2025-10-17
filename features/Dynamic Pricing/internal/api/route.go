package api

import (
	"net/http"

	"VOID/internal/services"
	"github.com/gin-gonic/gin"
)

func ComputeRoute(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")
	if start == "" || end == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start and end required"})
		return
	}
	path, cost := services.ComputeOptimalRoute(start, end)
	c.JSON(http.StatusOK, gin.H{"path": path, "cost": cost})
}
