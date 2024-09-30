package models

import (
	"time"

	utils "github.com/juliuscecilia33/sagev2/utility"
)

type Item struct {
	ID              	uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            	string    `json:"name" gorm:"type:varchar(255)"`
	Description     	string    `json:"description" gorm:"type:text"`
	ItemType			string	  `json:"itemType" gorm:"type:text"`
	UnlockConditions	utils.JSONMap `json:"unlockConditions" gorm:"type:jsonb"`
	LevelImages			utils.JSONMap `json:"levelImages" gorm:"type:jsonb"`
	CreatedAt       	time.Time	`json:"createdAt"`
	UpdatedAt       	time.Time	`json:"udpatedAt"`
}

