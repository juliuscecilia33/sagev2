package repositories

import (
	"context"

	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func (r *ItemRepository) GetMany(ctx context.Context) ([]*models.Item, error) {
	items := []*models.Item{}

	res := r.db.Model(&models.Item{}).Order("updated_at desc").Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}

	return items, nil
}

func (r *ItemRepository) GetOne(ctx context.Context, itemId uint) (*models.Item, error) {
	item := &models.Item{}

	res := r.db.Model(item).Where("id = ?", itemId).First(item)

	if res.Error != nil {
		return nil, res.Error
	}

	return item, nil
}

func (r *ItemRepository) CreateOne(ctx context.Context, item *models.Item) (*models.Item, error) {
	res := r.db.Model(item).Create(item)

	if res.Error != nil {
		return nil, res.Error
	}

	return item, nil
}

func (r *ItemRepository) UpdateOne(ctx context.Context, itemId uint, updateData map[string]interface{}) (*models.Item, error) {
	item := &models.Item{}

	updateRes := r.db.Model(item).Where("id = ?", itemId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", itemId).First(item)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return item, nil
}

func NewItemRepository(db *gorm.DB) models.ItemRepository {
	return &ItemRepository{
		db: db,
	}
}