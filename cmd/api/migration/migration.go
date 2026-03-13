package migration

import (
	"log"

	"github.com/AdityaAWP/IdiomaMate/internal/domain"
	"github.com/AdityaAWP/IdiomaMate/pkg/database"
)

func Migrate() {
	if err := database.DB.AutoMigrate(
		&domain.User{},
	); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("INFO: Database migrated")
}
