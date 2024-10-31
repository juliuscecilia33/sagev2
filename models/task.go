package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/utils"
)

type Task struct {
	ID 					uuid.UUID			`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RewardID			uuid.UUID			`json:"rewardId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Name				string				`json:"name" gorm:"type:text; not null"`
	Description			string				`json:"description" gorm:"type:text"`
	Requirements		utils.NestedJSONMap	`json:"requirements" gorm:"type:jsonb; not null"`
	Type				string				`json:"type" gorm:"type:text"`
	Reward      		Reward          	`gorm:"foreignKey:RewardID;references:ID"`
}

type TaskRepository interface {
	GetMany(ctx context.Context) ([]*Task, error)
	GetOne(ctx context.Context, taskId uuid.UUID) (*Task, error)
	CreateOne(ctx context.Context, task *Task) (*Task, error)
	UpdateOne(ctx context.Context, taskId uuid.UUID, updateData map[string]interface{}) (*Task, error)
	DeleteOne(ctx context.Context, taskId uuid.UUID) error
}