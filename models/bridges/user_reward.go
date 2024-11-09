package bridges

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
)

type UserReward struct {
	ID 				uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	RewardID		uuid.UUID			`json:"rewardId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	EarnedAt   		time.Time           `json:"started_at" gorm:"type:timestamp"`
	CreatedAt   	time.Time			`json:"created_at"`
	UpdatedAt   	time.Time			`json:"udpated_at"`
	User            models.User         `gorm:"foreignKey:UserID;references:ID"`
	Reward      	models.Reward       `gorm:"foreignKey:RewardID;references:ID"`
}


type UserRewardRepository interface {
	GetMany(ctx context.Context) ([]*UserReward, error)
	GetOne(ctx context.Context, userRewardId uuid.UUID) (*UserReward, error)
	CreateOne(ctx context.Context, userReward *UserReward) (*UserReward, error)
	UpdateOne(ctx context.Context, userRewardId uuid.UUID, updateData map[string]interface{}) (*UserReward, error)
	DeleteOne(ctx context.Context, userRewardId uuid.UUID) error
}