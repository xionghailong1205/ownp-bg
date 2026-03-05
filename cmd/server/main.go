package main

import (
	"context"
	"log"
	"time"

	"ownp-bg/internal/config"
	"ownp-bg/internal/database"
	"ownp-bg/internal/handler"
	"ownp-bg/internal/repository"
	"ownp-bg/internal/router"
	"ownp-bg/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if cfg.Env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	databaseURL, err := cfg.ActiveDatabaseURL()
	if err != nil {
		log.Fatalf("failed to get active database url: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.NewPostgresDB(ctx, databaseURL)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	timeRepo := repository.NewPostgresTimeRepository(db)
	healthService := service.NewHealthService(timeRepo)
	healthHandler := handler.NewHealthHandler(healthService)

	engine := router.New(healthHandler)
	if err := engine.Run(cfg.ServerAddr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
