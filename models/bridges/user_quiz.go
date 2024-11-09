package bridges

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
	"github.com/juliuscecilia33/sagev2/utils"
)

type UserQuiz struct {
	ID 				uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	QuizID         	uuid.UUID     		`json:"quizId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Progress		utils.NestedJSONMap `json:"progress" gorm:"type:jsonb"`
	Score			int              	`json:"score" gorm:"type:int"`
	StartedAt   	time.Time           `json:"started_at" gorm:"type:timestamp"` // Timestamp for when the quiz started
    CompletedAt 	time.Time           `json:"completed_at" gorm:"type:timestamp"` // Timestamp for when the quiz was completed
	CreatedAt   	time.Time			`json:"created_at"`
	UpdatedAt   	time.Time			`json:"udpated_at"`
	User            models.User          `gorm:"foreignKey:UserID;references:ID"`
	Quiz          	models.Quiz       	  `gorm:"foreignKey:QuizID;references:ID"`
}

type UserQuizRepository interface {
	GetMany(ctx context.Context) ([]*UserQuiz, error)
	GetOne(ctx context.Context, userQuizId uuid.UUID) (*UserQuiz, error)
	GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*UserQuiz, error)
	CreateOne(ctx context.Context, userQuiz *UserQuiz) (*UserQuiz, error)
	UpdateOne(ctx context.Context, userQuizId uuid.UUID, updateData map[string]interface{}) (*UserQuiz, error)
	DeleteOne(ctx context.Context, userQuizId uuid.UUID) error
}