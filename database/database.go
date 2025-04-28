package database

import (
	"os"
	"redis-caching/model"
	"redis-caching/seed"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	// connect to database
	databaseConfig := os.Getenv("DATABASE_CONFIG")
	db, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// migrate user
	db.AutoMigrate(&model.User{})

	// seed user data to user table
	seed.SeedUser(db)

	return db, err
}
