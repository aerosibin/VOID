package main

import (
	"fmt"
	"log"
	"os"

	"VOID/config"
	"VOID/internal/api"
	"VOID/internal/cache"
	"VOID/internal/db"
	"VOID/internal/ws"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// DB init (sqlite default; use DATABASE_URL env for Postgres)
	gormDB, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	if err := db.AutoMigrate(gormDB); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	// Cache (Redis optional)
	cacheClient := cache.NewCache(cfg)

	// WebSocket manager
	manager := ws.NewManager()
	go manager.Run()

	// Gin router
	r := gin.Default()
	api.RegisterRoutes(r, gormDB, cacheClient, manager, cfg)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("🚀 VOID backend listening on %s\n", addr)
	if err := r.Run(addr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
