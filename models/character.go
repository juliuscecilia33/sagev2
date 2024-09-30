package models

import (
	"context"
	"time"

	utils "github.com/juliuscecilia33/sagev2/utility"
)

type Character struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string    `json:"name" gorm:"type:varchar(255)"`
	Description     string    `json:"description" gorm:"type:text"`
	FruitMultipliers utils.JSONMap   `json:"fruitMultipliers" gorm:"type:jsonb"` // JSONB for PostgreSQL
	LevelImages      utils.JSONMap   `json:"levelIMages" gorm:"type:jsonb"` // JSONB for PostgreSQL
	CreatedAt       time.Time	`json:"createdAt"`
	UpdatedAt       time.Time	`json:"updatedAt"`
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId uint) (*Character, error)
	CreateOne(ctx context.Context, character *Character) (*Character, error)
	UpdateOne(ctx context.Context, characterId uint, updateData map[string] interface{}) (*Character, error)
	DeleteOne(ctx context.Context, characterId uint) error
}