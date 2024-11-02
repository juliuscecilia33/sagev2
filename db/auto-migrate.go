package db

import (
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	models := []interface{}{
		&models.Character{},
		&models.Item{},
		&models.User{},
		&models.UserKid{},
		&models.Quiz{},
		&models.Reward{},
		&models.Task{},
	}

	return db.AutoMigrate(models...)
}