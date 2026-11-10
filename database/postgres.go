package database

import (
	"log"

	"github.com/rajpatelbot/url-shortner/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This function will established connection with DB and return DB instance.
func NewDatabase() *gorm.DB {
	dsn := utils.GenerateDSN()

	// Open DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
