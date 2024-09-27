package repositories

import (
	"context"

	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func (r *CharacterRepository) GetMany(ctx context.Context) ([]*models.Character, error) {

	characters := []*models.Character{}

	res := r.db.Model(&models.Character{}).Order("updated_at desc").Find(&characters)

	if res.Error != nil {
		return nil, res.Error
	}

	return characters, nil
}


func (r *CharacterRepository) GetOne(ctx context.Context, characterId string) (*models.Character, error) {
	character := &models.Character{}

	res := r.db.Model(character).Where("id = ?", characterId).First(character)

	if res.Error != nil {
		return nil, res.Error
	}

	return character, nil
}

func (r *CharacterRepository) CreateOne(ctx context.Context, character *models.Character) (*models.Character, error) {
	res := r.db.Model(character).Create(character)

	if res.Error != nil {
		return nil, res.Error
	}

	return character, nil
}

func (r *CharacterRepository) UpdateOne(ctx context.Context, characterId uint, updateData map[string]interface{}) (*models.Character, error) {
	character := &models.Character{}

	updateRes := r.db.Model(character).Where("id = ?", characterId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", characterId).First(character)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return character, nil
}

func NewCharacterRepository(db *gorm.DB) models.CharacterRepository {
	return &CharacterRepository{
		db: db,
	}
}