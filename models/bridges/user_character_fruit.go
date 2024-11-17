package bridges

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
)

type UserCharacterFruit struct {
	ID 					uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          	uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	CharacterID			uuid.UUID			`json:"characterId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	UserCharacterId		uuid.UUID         	`json:"userCharacterId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	FruitName			string 				`json:"fruit_name" gorm:"type:varchar(255)"`
	Points				int					`json:"points" gorm:"type:int"`
	Multiplier			int					`json:"multiplier" gorm:"type:int"`
	CreatedAt   		time.Time			`json:"created_at"`
	UpdatedAt   		time.Time			`json:"updated_at"`
	User            	models.User         `gorm:"foreignKey:UserID;references:ID"`
	Character     		models.Character    `gorm:"foreignKey:CharacterID;references:ID"`
	UserCharacter 		UserCharacter 		`gorm:"foreignKey:UserCharacterId;references:ID"`
}

type UserCharacterFruitRepository interface {
	GetMany(ctx context.Context) ([]*UserCharacterFruit, error)
	GetAllByUserCharacter(ctx context.Context, userId uuid.UUID, characterId uuid.UUID) ([]*UserCharacterFruit, error)
	GetOne(ctx context.Context, userCharacterFruitId uuid.UUID) (*UserCharacterFruit, error)
	CreateOne(ctx context.Context, userCharacterFruit *UserCharacterFruit) (*UserCharacterFruit, error)
	UpdateOne(ctx context.Context, userCharacterFruitId uuid.UUID, updateData map[string]interface{}) (*UserCharacterFruit, error)
	DeleteOne(ctx context.Context, userCharacterFruitId uuid.UUID) error
}