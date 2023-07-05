package database

import (
	"log"

	"github.com/athunlal/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %s", err)
	}

	// Perform auto-migration to create tables
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to perform auto-migration: %s", err)
	}

	DB = db
}
