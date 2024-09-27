package repositories

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("Something went wrong!")
	}

	return characters, nil
}

func (r *CharacterRepository) GetOne(ctx context.Context, characterId string) (*models.Character, error) {
	return nil, nil
}

func (r *CharacterRepository) CreateOne(ctx context.Context, character models.Character) (*models.Character, error) {
	return nil, nil
}

func NewCharacterRepository(db *gorm.DB) models.CharacterRepository {
	return &CharacterRepository{
		db: db,
	}
}