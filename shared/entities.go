package shared

import (
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
	"github.com/juliuscecilia33/sagev2/utils"
)

type UserCharacter struct {
	ID 				uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	CharacterID		uuid.UUID			`json:"characterId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	ExperiencePoints	int         	`json:"experience_points" gorm:"type:int"`
	EquippedItems	utils.NestedJSONMap `json:"equipped_items" gorm:"type:jsonb"`
	Stats			utils.NestedJSONMap `json:"stats" gorm:"type:jsonb"`
	CreatedAt   	time.Time			`json:"created_at"`
	UpdatedAt   	time.Time			`json:"updated_at"`
	User            models.User         `gorm:"foreignKey:UserID;references:ID"`
	Character     	models.Character    `gorm:"foreignKey:CharacterID;references:ID"`
}