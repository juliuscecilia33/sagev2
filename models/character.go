package models

import (
	"context"
	"time"
)


type Character struct {
	ID              string            `gorm:"primaryKey"`    // Mark ID as primary key
	Name            string            `gorm:"type:varchar(255)"`
	Description     string            `gorm:"type:text"`
	FruitMultipliers map[string]string `gorm:"type:json"`     // Store map as JSON in DB
	LevelImages     map[string]string `gorm:"type:json"`     // Store map as JSON in DB
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId string) (*Character, error)
	CreateOne(ctx context.Context, character Character) (*Character, error)
}