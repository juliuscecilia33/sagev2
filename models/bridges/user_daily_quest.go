package bridges

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models" // import models
)

type UserDailyQuest struct {
	ID 				uuid.UUID     		`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID          uuid.UUID     		`json:"userId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	DailyQuestID	uuid.UUID			`json:"dailyQuestId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Completed		bool				`json:"completed" gorm:"type:boolean;default:false"`
	CompletedAt   	time.Time			`json:"completed_at"`
	CreatedAt   	time.Time			`json:"created_at"`
	UpdatedAt   	time.Time			`json:"udpated_at"`
	User            models.User         `gorm:"foreignKey:UserID;references:ID"`
	DailyQuest      models.DailyQuest    `gorm:"foreignKey:RewardID;references:ID"`
}


type UserDailyQuestRepository interface {
	GetMany(ctx context.Context) ([]*UserDailyQuest, error)
	GetOne(ctx context.Context, userDailyQuestId uuid.UUID) (*UserDailyQuest, error)
	CreateOne(ctx context.Context, UserDailyQuest *UserDailyQuest) (*UserDailyQuest, error)
	UpdateOne(ctx context.Context, userDailyQuestId uuid.UUID, updateData map[string]interface{}) (*UserDailyQuest, error)
	DeleteOne(ctx context.Context, userDailyQuestId uuid.UUID) error
}