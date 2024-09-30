package models

import (
	"context"
	"time"

	utils "github.com/juliuscecilia33/sagev2/utility"
)

type Character struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	FruitMultipliers utils.JSONMap   `gorm:"type:jsonb"` // JSONB for PostgreSQL
	LevelImages      utils.JSONMap   `gorm:"type:jsonb"` // JSONB for PostgreSQL
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId uint) (*Character, error)
	CreateOne(ctx context.Context, character *Character) (*Character, error)
	UpdateOne(ctx context.Context, characterId uint, updateData map[string] interface{}) (*Character, error)
	DeleteOne(ctx context.Context, characterId uint) error
}