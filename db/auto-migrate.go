package db

import (
	"github.com/juliuscecilia33/sagev2/models"
	"github.com/juliuscecilia33/sagev2/models/bridges"
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
		&models.DailyQuest{},
		&bridges.UserQuiz{},
		&bridges.UserReward{},
		&bridges.UserDailyQuest{},
		&bridges.UserCharacter{},
		&bridges.UserTask{},
	}

	return db.AutoMigrate(models...)
}