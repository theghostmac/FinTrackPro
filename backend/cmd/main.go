package main

import (
	"fintrackpro/backend/api"
	"fintrackpro/backend/config"
	"fintrackpro/backend/internal/database"
	"fintrackpro/backend/internal/models"
	"go.uber.org/zap"
	"strings"
)

func main() {
	// Initialize the logger.
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			zap.L().Sugar().Errorf("Failed to sync: %e", err)
		}
	}(logger) // Flushes buffer, if any

	// Database connection.
	dbConnection, err := database.NewDBConnection()
	if err != nil {
		zap.L().Sugar().Errorf("Failed to initialize db connection: %v", err)
		return
	}

	// Automigrate the models.
	err = dbConnection.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.BudgetPlan{},
	)
	if err != nil {
		zap.L().Sugar().Errorf("Failed to auto-migrate database models: %v", err)
		return
	}

	// Load configuration.
	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Sugar().Errorf("Failed to load configuration: %v", err)
		return
	}

	// Initialize the router (assuming InitRouter is defined elsewhere).
	router := api.InitRouter()

	// Use the configuration's listen address if available, otherwise default to ":8080".
	listenAddr := cfg.ListenAddress
	if listenAddr == "" {
		listenAddr = ":8080"
	} else if !strings.HasPrefix(listenAddr, ":") {
		listenAddr = ":" + listenAddr
	}

	zap.L().Sugar().Infof("Starting server at: %v", listenAddr)

	// Start the Gin server.
	if err := router.Run(listenAddr); err != nil {
		zap.L().Sugar().Errorf("Failed to run server: %v", err)
	}
}
