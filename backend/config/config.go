package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()

type Config struct {
	ListenAddress string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}

	return &Config{
		ListenAddress: os.Getenv("SERVER_PORT"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}, nil
}
