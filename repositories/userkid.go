package repositories

import (
	"context"

	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type UserKidRepository struct {
	db *gorm.DB
}

func (r *UserKidRepository) GetAllParentKids(ctx context.Context, parentId uint) ([]*models.UserKid, error) {
	kids := []*models.UserKid{}

	res := r.db.Model(&models.UserKid{}).Where("parent_id = ?", parentId).Preload("User").Find(&kids)

	if res.Error != nil {
		return nil, res.Error
	}

	return kids, nil
}

func (r *UserKidRepository) GetOne(ctx context.Context, userKidId uint) (*models.UserKid, error) {
	kid := &models.UserKid{}

	res := r.db.Model(kid).Where("id = ?", userKidId).Preload("User").First(kid)

	if res.Error != nil {
		return nil, res.Error
	}

	return kid, nil
}

func (r *UserKidRepository) CreateOne(ctx context.Context, kid *models.UserKid) (*models.UserKid, error) {
	res := r.db.Model(kid).Create(kid)

	if res.Error != nil {
		return nil, res.Error
	}

	return kid, nil
}

func (r *UserKidRepository) UpdateOne(ctx context.Context, userKidId uint, updateData map[string]interface{}) (*models.UserKid, error) {
	kid := &models.UserKid{}

	updateRes := r.db.Model(kid).Where("id = ?", userKidId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userKidId).First(kid)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return kid, nil
}

func NewUserKidRepository(db *gorm.DB) models.UserKidRepository {
	return &UserKidRepository{
		db: db,
	}
}