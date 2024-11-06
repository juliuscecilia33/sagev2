package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type DailyQuestRepository struct {
	db *gorm.DB
}

func (r *DailyQuestRepository) GetMany(ctx context.Context) ([]*models.DailyQuest, error) {
	dailyQuests := []*models.DailyQuest{}

	res := r.db.Model(&models.DailyQuest{}).Preload("Reward").Order("quest_date desc").Find(&dailyQuests)

	if res.Error != nil {
		return nil, res.Error
	}

	return dailyQuests, nil
}

func (r *DailyQuestRepository) GetOne(ctx context.Context, dailyQuestId uuid.UUID) (*models.DailyQuest, error) {
	dailyQuest := &models.DailyQuest{}

	res := r.db.Model(dailyQuest).Preload("Reward").Where("id = ?", dailyQuestId).First(dailyQuest)

	if res.Error != nil {
		return nil, res.Error
	}

	return dailyQuest, nil
}

func (r *DailyQuestRepository) CreateOne(ctx context.Context, dailyQuest *models.DailyQuest) (*models.DailyQuest, error) {
	res := r.db.Create(dailyQuest)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("Reward").First(dailyQuest, dailyQuest.ID).Error; err != nil {
		return nil, err
	}

	return dailyQuest, nil
}

func (r DailyQuestRepository) UpdateOne(ctx context.Context, dailyQuestId uuid.UUID, updateData map[string]interface{}) (*models.DailyQuest, error) {
	dailyQuest := &models.DailyQuest{}

	updateRes := r.db.Model(dailyQuest).Where("id = ?", dailyQuestId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", dailyQuestId).First(dailyQuest)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return dailyQuest, nil
}

func (r *DailyQuestRepository) DeleteOne(ctx context.Context, dailyQuestId uuid.UUID) error {
	res := r.db.Delete(&models.DailyQuest{}, dailyQuestId)

	return res.Error
}


func NewDailyQuestRepository(db *gorm.DB) models.DailyQuestRepository {
	return &DailyQuestRepository{
		db: db,
	}
}