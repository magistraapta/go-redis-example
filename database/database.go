package database

import (
	"os"
	"redis-caching/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	databaseConfig := os.Getenv("DATABASE_CONFIG")
	db, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	return db, err
}
