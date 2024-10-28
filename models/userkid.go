package models

import (
	"context"
	"time"

	"github.com/juliuscecilia33/sagev2/utils"
)

type UserKid struct {
	ID        			uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    			uint `json:"userId" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"` // Kid's user ID
	ParentID  			uint `json:"parentId" gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"` // Parent's user ID
	CurrentFruitStats	utils.JSONMap `json:"current_fruit_stats" gorm:"type:jsonb"`
	BibleProgress		utils.JSONMap `json:"bible_progress" gorm:"type:jsonb"`
	CreatedAt       	time.Time	`json:"created_at"`
	UpdatedAt       	time.Time	`json:"udpated_at"`
	User      			User `gorm:"foreignKey:UserID;references:ID"` // Establish relation to User for UserID
	Parent    			User `gorm:"foreignKey:ParentID;references:ID"` // Establish relation to User for ParentID
}

type UserKidRepository interface {
	GetAllParentKids(ctx context.Context, parentId uint) ([]*UserKid, error)
	GetOne(ctx context.Context, userKidId uint) (*UserKid, error)
	CreateOne(ctx context.Context, userKid *UserKid) (*UserKid, error)
	UpdateOne(ctx context.Context, userKidId uint, updateData map[string]interface{}) (*UserKid, error)
}
