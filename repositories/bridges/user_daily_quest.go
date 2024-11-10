package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserDailyQuestRepository struct {
	db *gorm.DB
}

func (r *UserDailyQuestRepository) GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*bridges.UserDailyQuest, error) {
	specific_user_daily_quests := []*bridges.UserDailyQuest{}

	res := r.db.Model(&bridges.UserDailyQuest{}).Where("user_id = ?", userId).Preload("User").Preload("DailyQuest").Find(&specific_user_daily_quests)

	if res.Error != nil {
		return nil, res.Error
	}

	return specific_user_daily_quests, nil
}

func (r *UserDailyQuestRepository) GetMany(ctx context.Context) ([]*bridges.UserDailyQuest, error) {
	user_daily_quests := []*bridges.UserDailyQuest{}

	res := r.db.Model(&bridges.UserDailyQuest{}).Preload("User").Preload("DailyQuest").Order("updated_at desc").Find(&user_daily_quests)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_daily_quests, nil
}

func (r *UserDailyQuestRepository) GetOne(ctx context.Context, userDailyQuestId uuid.UUID) (*bridges.UserDailyQuest, error) {
	user_daily_quest := &bridges.UserDailyQuest{}

	res := r.db.Model(user_daily_quest).Preload("User").Preload("Reward").Where("id = ?", userDailyQuestId).First(user_daily_quest)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_daily_quest, nil
}

func (r *UserDailyQuestRepository) CreateOne(ctx context.Context, user_daily_quest *bridges.UserDailyQuest) (*bridges.UserDailyQuest, error) {
	res := r.db.Create(user_daily_quest)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("DailyQuest").Preload("DailyQuest.Reward").First(user_daily_quest, user_daily_quest.ID).Error; err != nil {
		return nil, err
	}

	return user_daily_quest, nil
}

func (r UserDailyQuestRepository) UpdateOne(ctx context.Context, userDailyQuestId uuid.UUID, updateData map[string]interface{}) (*bridges.UserDailyQuest, error) {
	user_daily_quest := &bridges.UserDailyQuest{}

	updateRes := r.db.Model(user_daily_quest).Where("id = ?", userDailyQuestId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userDailyQuestId).First(user_daily_quest)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_daily_quest, nil
}

func (r *UserDailyQuestRepository) DeleteOne(ctx context.Context, userDailyQuestId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserDailyQuest{}, userDailyQuestId)

	return res.Error
}

func NewUserDailyQuestRepository(db *gorm.DB) bridges.UserDailyQuestRepository {
	return &UserDailyQuestRepository{
		db: db,
	}
}