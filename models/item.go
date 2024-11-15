package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	utils "github.com/juliuscecilia33/sagev2/utils"
)

type Item struct {
	ID         uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name            	string    `json:"name" gorm:"type:varchar(255)"`
	Description     	string    `json:"description" gorm:"type:text"`
	ItemType			string	  `json:"item_type" gorm:"type:text"`
	UnlockConditions	utils.JSONMap `json:"unlock_conditions" gorm:"type:jsonb"`
	LevelImages			utils.JSONMap `json:"level_images" gorm:"type:jsonb"`
	CreatedAt       	time.Time	`json:"created_at"`
	UpdatedAt       	time.Time	`json:"updated_at"`
}

type ItemRepository interface {
	GetMany(ctx context.Context) ([]*Item, error)
	GetOne(ctx context.Context, itemId uuid.UUID) (*Item, error)
	CreateOne(ctx context.Context, item *Item) (*Item, error)
	UpdateOne(ctx context.Context, itemId uuid.UUID, updateData map[string]interface{}) (*Item, error)
}

