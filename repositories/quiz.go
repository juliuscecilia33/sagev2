package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func (r *QuizRepository) GetMany(ctx context.Context) ([]*models.Quiz, error) {
	quizzes := []*models.Quiz{}

	res := r.db.Model(&models.Quiz{}).Order("updated_at desc").Find(&quizzes)

	if res.Error != nil {
		return nil, res.Error
	}

	return quizzes, nil
}


func (r *QuizRepository) GetOne(ctx context.Context, quizId uuid.UUID) (*models.Quiz, error) {
	quiz := &models.Quiz{}

	res := r.db.Model(quiz).Where("id = ?", quizId).First(quiz)

	if res.Error != nil {
		return nil, res.Error
	}

	return quiz, nil
}

func (r *QuizRepository) CreateOne(ctx context.Context, quiz *models.Quiz) (*models.Quiz, error) {
	res := r.db.Model(quiz).Create(quiz)

	if res.Error != nil {
		return nil, res.Error
	}

	return quiz, nil
}

func (r *QuizRepository) UpdateOne(ctx context.Context, quizId uuid.UUID, updateData map[string]interface{}) (*models.Quiz, error) {
	quiz := &models.Quiz{}

	updateRes := r.db.Model(quiz).Where("id = ?", quizId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", quizId).First(quiz)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return quiz, nil
}

func NewQuizRepository(db *gorm.DB) models.QuizRepository {
	return &QuizRepository{
		db: db,
	}
}