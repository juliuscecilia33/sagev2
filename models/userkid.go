package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/utils"
)

type UserKid struct {
	ID               uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID           uuid.UUID     `json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"` // Kid's user ID
	ParentID         uuid.UUID     `json:"parentId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"` // Parent's user ID
	CurrentFruitStats utils.JSONMap `json:"current_fruit_stats" gorm:"type:jsonb"`
	BibleProgress    utils.JSONMap `json:"bible_progress" gorm:"type:jsonb"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	User             User          `gorm:"foreignKey:UserID;references:ID"` // Establish relation to User for UserID
	Parent           User          `gorm:"foreignKey:ParentID;references:ID"` // Establish relation to User for ParentID
}

type UserKidRepository interface {
	GetAllParentKids(ctx context.Context, parentId uuid.UUID) ([]*UserKid, error)
	GetOne(ctx context.Context, userKidId uuid.UUID) (*UserKid, error)
	CreateOne(ctx context.Context, userKid *UserKid) (*UserKid, error)
	UpdateOne(ctx context.Context, userKidId uuid.UUID, updateData map[string]interface{}) (*UserKid, error)
	DeleteOne(ctx context.Context, userKidId uuid.UUID) error
}
