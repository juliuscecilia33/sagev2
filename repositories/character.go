package repositories

import (
	"context"

	"github.com/juliuscecilia33/sagev2/models"
)

type CharacterRepository struct {
	db any
}

func (r *CharacterRepository) GetMany(ctx context.Context) ([]*models.Character, error) {
	return nil, nil
}

func (r *CharacterRepository) GetOne(ctx context.Context, characterId string) (*models.Character, error) {
	return nil, nil
}

func (r *CharacterRepository) CreateOne(ctx context.Context, character models.Character) (*models.Character, error) {
	return nil, nil
}

func NewCharacterRepository(db any) models.CharacterRepository {
	return &CharacterRepository{
		db: db,
	}
}