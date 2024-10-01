package db

import (
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Character{}, &models.Item{})
}