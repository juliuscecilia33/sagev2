package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	utils "github.com/juliuscecilia33/sagev2/utils"
)

type Quiz struct {
	ID         	uuid.UUID			`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title      	string      		`json:"title" gorm:"type:text"`
	Description string				`json:"description" gorm:"type:text"`
	Content		utils.NestedJSONMap	`json:"content" gorm:"type:jsonb"`
	CreatedAt   time.Time			`json:"created_at"`
	UpdatedAt   time.Time			`json:"updated_at"`
}

type QuizRepository interface {
	GetMany(ctx context.Context) ([]*Quiz, error)
	GetOne(ctx context.Context, quizId uuid.UUID) (*Quiz, error)
	CreateOne(ctx context.Context, quiz *Quiz) (*Quiz, error)
	UpdateOne(ctx context.Context, quizId uuid.UUID, updateData map[string]interface{}) (*Quiz, error)
	DeleteOne(ctx context.Context, quizId uuid.UUID) error
}

