package repositories

import (
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type ItemRepository struct {
	
}

func NewItemRepository(db *gorm.DB) models.CharacterRepository {
	return &ItemRepository{
		db: db,
	}
}