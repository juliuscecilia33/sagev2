package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/utils"
)

type Reward struct {
	ID 			uuid.UUID			`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name		string      		`json:"name" gorm:"type:text"`
	Description string				`json:"description" gorm:"type:text; not null"`
	Type		string				`json:"type" gorm:"type:text; not null"`
	Details		utils.NestedJSONMap	`json:"details" gorm:"type:jsonb"`
	CreatedAt 	time.Time			`json:"created_at"`
	UpdatedAt 	time.Time			`json:"udpated_at"`
}

type RewardRepository interface {
	GetMany(ctx context.Context) ([]*Reward, error)
	GetOne(ctx context.Context, rewardId uuid.UUID) (*Reward, error)
	CreateOne(ctx context.Context, reward *Reward) (*Reward, error)
	UpdateOne(ctx context.Context, rewardId uuid.UUID, updateData map[string]interface{}) (*Reward, error)
	DeleteOne(ctx context.Context, rewardId uuid.UUID) error
}
