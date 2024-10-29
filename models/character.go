package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	utils "github.com/juliuscecilia33/sagev2/utils"
)

type Character struct {
	ID         uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name            string    `json:"name" gorm:"type:varchar(255)"`
	Description     string    `json:"description" gorm:"type:text"`
	FruitMultipliers utils.JSONMap   `json:"fruit_multipliers" gorm:"type:jsonb"` // JSONB for PostgreSQL
	LevelImages      utils.JSONMap   `json:"level_images" gorm:"type:jsonb"` // JSONB for PostgreSQL
	CreatedAt       time.Time	`json:"created_at"`
	UpdatedAt       time.Time	`json:"updated_at"`
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId uuid.UUID) (*Character, error)
	CreateOne(ctx context.Context, character *Character) (*Character, error)
	UpdateOne(ctx context.Context, characterId uuid.UUID, updateData map[string] interface{}) (*Character, error)
	DeleteOne(ctx context.Context, characterId uuid.UUID) error
}