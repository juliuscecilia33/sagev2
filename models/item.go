package models

import (
	"time"

	utils "github.com/juliuscecilia33/sagev2/utility"
)

type Item struct {
	ID              	uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            	string    `gorm:"type:varchar(255)"`
	Description     	string    `gorm:"type:text"`
	ItemType			string	  `gorm:"type:text"`
	UnlockConditions	utils.JSONMap `gorm:"type:jsonb"`
	LevelImages			utils.JSONMap `gorm:"type:jsonb"`
	CreatedAt       	time.Time
	UpdatedAt       	time.Time
}

