package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func (r *RewardRepository) GetMany(ctx context.Context) ([]*models.Reward, error) {
	rewards := []*models.Reward{}

	res := r.db.Model(&models.Reward{}).Order("updated_at desc").Find(&rewards)

	if res.Error != nil {
		return nil, res.Error
	}

	return rewards, nil
}

func (r *RewardRepository) GetOne(ctx context.Context, rewardId uuid.UUID) (*models.Reward, error) {
	reward := &models.Reward{}

	res := r.db.Model(reward).Where("id = ?", rewardId).First(reward)

	if res.Error != nil {
		return nil, res.Error
	}

	return reward, nil
}

func (r *RewardRepository) CreateOne(ctx context.Context, reward *models.Reward) (*models.Reward, error) {
	res := r.db.Model(reward).Create(reward)

	if res.Error != nil {
		return nil, res.Error
	}

	return reward, nil
}

func (r RewardRepository) UpdateOne(ctx context.Context, rewardId uuid.UUID, updateData map[string]interface{}) (*models.Reward, error) {
	reward := &models.Reward{}

	updateRes := r.db.Model(reward).Where("id = ?", rewardId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", rewardId).First(reward)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return reward, nil
}

func (r *RewardRepository) DeleteOne(ctx context.Context, rewardId uuid.UUID) error {
	res := r.db.Delete(&models.Reward{}, rewardId)

	return res.Error
}

func NewRewardRepository(db *gorm.DB) models.RewardRepository {
	return &RewardRepository{
		db: db,
	}
}