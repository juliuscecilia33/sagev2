package bridges

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
)

type UserTask struct {
	ID 				uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	TaskID			uuid.UUID     		`json:"taskId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Status 			string 				`json:"status" gorm:"type:varchar(255);default:'Not started'"`
	CreatedAt   	time.Time			`json:"created_at"`
	UpdatedAt   	time.Time			`json:"udpated_at"`
	User            models.User         `gorm:"foreignKey:UserID;references:ID"`
	Task     		models.Task    		`gorm:"foreignKey:TaskID;references:ID"`
}

type UserTaskRepository interface {
	GetMany(ctx context.Context) ([]*UserTask, error)
	GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*UserTask, error)
	GetOne(ctx context.Context, userTaskId uuid.UUID) (*UserTask, error)
	CreateOne(ctx context.Context, userTask *UserTask) (*UserTask, error)
	UpdateOne(ctx context.Context, userTaskId uuid.UUID, updateData map[string]interface{}) (*UserTask, error)
	DeleteOne(ctx context.Context, userTaskId uuid.UUID) error
}