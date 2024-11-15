package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/utils"
)

type DailyQuest struct {
	ID 					uuid.UUID			`json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RewardID			uuid.UUID			`json:"rewardId" gorm:"type:uuid;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Name				string				`json:"name" gorm:"type:text; not null"`
	Description			string				`json:"description" gorm:"type:text"`
	Requirements		utils.NestedJSONMap	`json:"requirements" gorm:"type:jsonb; not null"`
	Type				string				`json:"type" gorm:"type:text"`
	QuestDate			string				`json:"quest_date" gorm:"type:text"`
	Reward      		Reward          	`gorm:"foreignKey:RewardID;references:ID"`
	CreatedAt 			time.Time			`json:"created_at"`
	UpdatedAt 			time.Time			`json:"updated_at"`
}

// QuestDate Example: "2024-11-05"

type DailyQuestRepository interface {
	GetMany(ctx context.Context) ([]*DailyQuest, error)
	GetOne(ctx context.Context, dailyQuestId uuid.UUID) (*DailyQuest, error)
	CreateOne(ctx context.Context, dailyQuest *DailyQuest) (*DailyQuest, error)
	UpdateOne(ctx context.Context, dailyQuestId uuid.UUID, updateData map[string]interface{}) (*DailyQuest, error)
	DeleteOne(ctx context.Context, dailyQuestId uuid.UUID) error
	GetByDate(ctx context.Context, questDate string) ([]*DailyQuest, error)
}