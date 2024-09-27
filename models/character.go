package models

import (
	"context"
	"time"
)


type Character struct {
	ID              uint            	`json:"id" gorm:"primaryKey;autoIncrement"`    // Mark ID as primary key
	Name            string            	`json:"name" gorm:"type:varchar(255)"`
	Description     string            	`json:"description" gorm:"type:text"`
	FruitMultipliers map[string]string 	`json:"fruitMultipliers" gorm:"type:json"`     // Store map as JSON in DB
	LevelImages     map[string]string 	`json:"levelImages" gorm:"type:json"`     // Store map as JSON in DB
	CreatedAt       time.Time			`json:"createdAt"`
	UpdatedAt       time.Time			`json:"updatedAt"`
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId uint) (*Character, error)
	CreateOne(ctx context.Context, character *Character) (*Character, error)
	UpdateOne(ctx context.Context, characterId uint, updateData map[string] interface{}) (*Character, error)
	DeleteOne(ctx context.Context, characterId uint) error
}