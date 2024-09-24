package models

import (
	"context"
	"time"
)


type Character struct {
	ID string
	Name string
	Description string
	FruitMultipliers map[string]string
	LevelImages map[string]string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CharacterRepository interface {
	GetMany(ctx context.Context) ([]*Character, error)
	GetOne(ctx context.Context, characterId string) (*Character, error)
	CreateOne(ctx context.Context, character Character) (*Character, error)
}