package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type UserKidRepository struct {
	db *gorm.DB
}

func (r *UserKidRepository) GetAllParentKids(ctx context.Context, parentId uuid.UUID) ([]*models.UserKid, error) {
	kids := []*models.UserKid{}

	res := r.db.Model(&models.UserKid{}).Where("parent_id = ?", parentId).Preload("User").Preload("Parent").Find(&kids)

	if res.Error != nil {
		return nil, res.Error
	}

	return kids, nil
}

func (r *UserKidRepository) GetOne(ctx context.Context, userKidId uuid.UUID) (*models.UserKid, error) {
	kid := &models.UserKid{}

	// Preload both User and Parent associations
	res := r.db.Where("id = ?", userKidId).
		Preload("User").
		Preload("Parent").
		First(kid)

	if res.Error != nil {
		return nil, res.Error
	}

	return kid, nil
}

func (r *UserKidRepository) CreateOne(ctx context.Context, kid *models.UserKid) (*models.UserKid, error) {
	res := r.db.Create(kid)
	if res.Error != nil {
		return nil, res.Error
	}

	// Load the User and Parent details after creation
	if err := r.db.Preload("User").Preload("Parent").First(kid, kid.ID).Error; err != nil {
		return nil, err
	}

	return kid, nil
}

func (r *UserKidRepository) UpdateOne(ctx context.Context, userKidId uuid.UUID, updateData map[string]interface{}) (*models.UserKid, error) {
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

func (r *UserKidRepository) DeleteOne(ctx context.Context, userKidId uuid.UUID) error {
	res := r.db.Delete(&models.UserKid{}, userKidId)

	return res.Error
}


func NewUserKidRepository(db *gorm.DB) models.UserKidRepository {
	return &UserKidRepository{
		db: db,
	}
}