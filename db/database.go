package db

import (
	"fmt"
	"log"

	"github.com/juliuscecilia33/sagev2/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s port=5432",
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)  // Fixing the formatting of the error message
	}

	log.Println("Connected to the database!")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Error running database migrations: %v", err)  // Fixing the formatting of the error message
	}

	return db
}
