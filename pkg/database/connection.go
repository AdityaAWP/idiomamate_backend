package database

import (
	"fmt"
	"log"

	"github.com/AdityaAWP/IdiomaMate/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	db := cfg.Database

	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
	var err error
	DB, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("INFO: Database initialized")

}
