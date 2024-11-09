package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserQuizRepository struct {
	db *gorm.DB
}

func (r *UserQuizRepository) GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*bridges.UserQuiz, error) {
	specific_user_quizzes := []*bridges.UserQuiz{}

	res := r.db.Model(&bridges.UserQuiz{}).Where("user_id = ?", userId).Preload("User").Find(&specific_user_quizzes)

	if res.Error != nil {
		return nil, res.Error
	}

	return specific_user_quizzes, nil
}

func (r *UserQuizRepository) GetMany(ctx context.Context) ([]*bridges.UserQuiz, error) {
	user_quizzes := []*bridges.UserQuiz{}

	res := r.db.Model(&bridges.UserQuiz{}).Preload("User").Preload("Quiz").Order("updated_at desc").Find(&user_quizzes)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_quizzes, nil
}

func (r *UserQuizRepository) GetOne(ctx context.Context, userQuizId uuid.UUID) (*bridges.UserQuiz, error) {
	user_quiz := &bridges.UserQuiz{}

	res := r.db.Model(user_quiz).Preload("User").Preload("Quiz").Where("id = ?", userQuizId).First(user_quiz)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_quiz, nil
}

func (r *UserQuizRepository) CreateOne(ctx context.Context, user_quiz *bridges.UserQuiz) (*bridges.UserQuiz, error) {
	res := r.db.Create(user_quiz)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("Quiz").First(user_quiz, user_quiz.ID).Error; err != nil {
		return nil, err
	}

	return user_quiz, nil
}

func (r UserQuizRepository) UpdateOne(ctx context.Context, userQuizId uuid.UUID, updateData map[string]interface{}) (*bridges.UserQuiz, error) {
	user_quiz := &bridges.UserQuiz{}

	updateRes := r.db.Model(user_quiz).Where("id = ?", userQuizId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userQuizId).First(user_quiz)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_quiz, nil
}

func (r *UserQuizRepository) DeleteOne(ctx context.Context, userQuizId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserQuiz{}, userQuizId)

	return res.Error
}

func NewUserQuizRepository(db *gorm.DB) bridges.UserQuizRepository {
	return &UserQuizRepository{
		db: db,
	}
}
