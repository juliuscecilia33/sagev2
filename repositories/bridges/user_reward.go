package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserRewardRepository struct {
	db *gorm.DB
}

func (r *UserRewardRepository) GetMany(ctx context.Context) ([]*bridges.UserReward, error) {
	user_rewards := []*bridges.UserReward{}

	res := r.db.Model(&bridges.UserReward{}).Preload("User").Preload("Reward").Order("updated_at desc").Find(&user_rewards)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_rewards, nil
}

func (r *UserRewardRepository) GetOne(ctx context.Context, userRewardId uuid.UUID) (*bridges.UserReward, error) {
	user_reward := &bridges.UserReward{}

	res := r.db.Model(user_reward).Preload("User").Preload("Reward").Where("id = ?", userRewardId).First(user_reward)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_reward, nil
}

func (r *UserRewardRepository) CreateOne(ctx context.Context, user_reward *bridges.UserReward) (*bridges.UserReward, error) {
	res := r.db.Create(user_reward)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("Reward").First(user_reward, user_reward.ID).Error; err != nil {
		return nil, err
	}

	return user_reward, nil
}

func (r UserRewardRepository) UpdateOne(ctx context.Context, userRewardId uuid.UUID, updateData map[string]interface{}) (*bridges.UserReward, error) {
	user_reward := &bridges.UserReward{}

	updateRes := r.db.Model(user_reward).Where("id = ?", userRewardId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userRewardId).First(user_reward)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_reward, nil
}

func (r *UserRewardRepository) DeleteOne(ctx context.Context, userRewardId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserReward{}, userRewardId)

	return res.Error
}

func NewUserRewardRepository(db *gorm.DB) bridges.UserRewardeRepository {
	return &UserRewardRepository{
		db: db,
	}
}