package repositories

import (
	"context"
	"time"

	"github.com/juliuscecilia33/sagev2/models"
)

type CharacterRepository struct {
	db any
}

func (r *CharacterRepository) GetMany(ctx context.Context) ([]*models.Character, error) {
	characters := []*models.Character{}

	characters = append(characters, &models.Character{
		ID:   "123123213891284918248917324912",
		Name: "Test Character",
		Description: "Hello this is Character Description",
		FruitMultipliers: map[string]string{
            "peace": "1.4",
            "gentleness": "1.3",
        },
		LevelImages: map[string]string{
            "levelOne": "www.google.com",
            "levelTwo": "www.google.com",
        },
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return characters, nil
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