package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserCharacterFruitRepository struct {
	db *gorm.DB
}

func (r *UserCharacterFruitRepository) GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*bridges.UserCharacterFruit, error) {
	specific_user_character_fruits := []*bridges.UserCharacterFruit{}

	res := r.db.Model(&bridges.UserCharacterFruit{}).Where("user_id = ?", userId).Preload("User").Preload("Character").Preload("UserCharacter").Find(&specific_user_character_fruits)

	if res.Error != nil {
		return nil, res.Error
	}

	return specific_user_character_fruits, nil
}

func (r *UserCharacterFruitRepository) GetMany(ctx context.Context) ([]*bridges.UserCharacterFruit, error) {
	user_character_fruits := []*bridges.UserCharacterFruit{}

	res := r.db.Model(&bridges.UserCharacterFruit{}).Preload("User").Preload("Character").Preload("UserCharacter").Order("updated_at desc").Find(&user_character_fruits)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_character_fruits, nil
}

func (r *UserCharacterFruitRepository) GetOne(ctx context.Context, userCharacterId uuid.UUID) (*bridges.UserCharacterFruit, error) {
	user_character_fruit := &bridges.UserCharacterFruit{}

	res := r.db.Model(user_character_fruit).Preload("User").Preload("Character").Preload("UserCharacter").Where("id = ?", userCharacterId).First(user_character_fruit)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_character_fruit, nil
}

func (r *UserCharacterFruitRepository) CreateOne(ctx context.Context, user_character_fruit *bridges.UserCharacterFruit) (*bridges.UserCharacterFruit, error) {
	res := r.db.Create(user_character_fruit)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("Character").Preload("UserCharacterFruit").First(user_character_fruit, user_character_fruit.ID).Error; err != nil {
		return nil, err
	}

	return user_character_fruit, nil
}

func (r UserCharacterFruitRepository) UpdateOne(ctx context.Context, userCharacterFruitId uuid.UUID, updateData map[string]interface{}) (*bridges.UserCharacterFruit, error) {
	user_character_fruit := &bridges.UserCharacterFruit{}

	updateRes := r.db.Model(user_character_fruit).Where("id = ?", userCharacterFruitId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userCharacterFruitId).First(user_character_fruit)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_character_fruit, nil
}

func (r *UserCharacterFruitRepository) DeleteOne(ctx context.Context, userCharacterFruitId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserCharacterFruit{}, userCharacterFruitId)

	return res.Error
}

func NewUserCharacterFruitRepository(db *gorm.DB) bridges.UserCharacterFruitRepository {
	return &UserCharacterFruitRepository{
		db: db,
	}
}
