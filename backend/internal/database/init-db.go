package database

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=fintrackpro_user password=f1ntr4ckpr0 dbname=fintrackpro_db port=5432 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Sugar().Errorf("Failed to open connection to db: %v", err)
		return nil, err
	}
	DB = db
	return db, nil
}
