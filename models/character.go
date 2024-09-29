package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type JSONMap map[string]string

// Custom MarshalJSON and UnmarshalJSON methods to handle the map serialization
func (m JSONMap) Value() (driver.Value, error) {
	// Marshal the map into a JSON string for the database
	return json.Marshal(m)
}

func (m *JSONMap) Scan(value interface{}) error {
	// Scan the JSON string from the database back into the map
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, m)
}

type Character struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	FruitMultipliers JSONMap   `gorm:"type:jsonb"` // JSONB for PostgreSQL
	LevelImages     JSONMap   `gorm:"type:jsonb"` // JSONB for PostgreSQL
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