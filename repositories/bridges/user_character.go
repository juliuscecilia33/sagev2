package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserCharacterRepository struct {
	db *gorm.DB
}

func (r *UserCharacterRepository) GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*bridges.UserCharacter, error) {
	specific_user_characters := []*bridges.UserCharacter{}

	res := r.db.Model(&bridges.UserCharacter{}).Where("user_id = ?", userId).Preload("User").Preload("Character").Find(&specific_user_characters)

	if res.Error != nil {
		return nil, res.Error
	}

	return specific_user_characters, nil
}

func (r *UserCharacterRepository) GetMany(ctx context.Context) ([]*bridges.UserCharacter, error) {
	user_characters := []*bridges.UserCharacter{}

	res := r.db.Model(&bridges.UserCharacter{}).Preload("User").Preload("Character").Order("updated_at desc").Find(&user_characters)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_characters, nil
}

func (r *UserCharacterRepository) GetOne(ctx context.Context, userCharacterId uuid.UUID) (*bridges.UserCharacter, error) {
	user_character := &bridges.UserCharacter{}

	res := r.db.Model(user_character).Preload("User").Preload("Character").Where("id = ?", userCharacterId).First(user_character)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_character, nil
}

func (r *UserCharacterRepository) CreateOne(ctx context.Context, user_character *bridges.UserCharacter) (*bridges.UserCharacter, error) {
	res := r.db.Create(user_character)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("Character").First(user_character, user_character.ID).Error; err != nil {
		return nil, err
	}

	return user_character, nil
}

func (r UserCharacterRepository) UpdateOne(ctx context.Context, userCharacterId uuid.UUID, updateData map[string]interface{}) (*bridges.UserCharacter, error) {
	user_character := &bridges.UserCharacter{}

	updateRes := r.db.Model(user_character).Where("id = ?", userCharacterId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userCharacterId).First(user_character)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_character, nil
}

func (r *UserCharacterRepository) DeleteOne(ctx context.Context, userCharacterId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserCharacter{}, userCharacterId)

	return res.Error
}

func NewUserCharacterRepository(db *gorm.DB) bridges.UserCharacterRepository {
	return &UserCharacterRepository{
		db: db,
	}
}

