package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func (r *TaskRepository) GetMany(ctx context.Context) ([]*models.Task, error) {
	tasks := []*models.Task{}

	res := r.db.Model(&models.Task{}).Preload("Reward").Order("updated_at desc").Find(&tasks)

	if res.Error != nil {
		return nil, res.Error
	}

	return tasks, nil
}

func (r *TaskRepository) GetOne(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	task := &models.Task{}

	res := r.db.Model(task).Preload("Reward").Where("id = ?", taskId).First(task)

	if res.Error != nil {
		return nil, res.Error
	}

	return task, nil
}

func (r *TaskRepository) CreateOne(ctx context.Context, task *models.Task) (*models.Task, error) {
	// res := r.db.Preload("Reward").Create(task)
	res := r.db.Create(task)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("Reward").First(task, task.ID).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (r TaskRepository) UpdateOne(ctx context.Context, taskId uuid.UUID, updateData map[string]interface{}) (*models.Task, error) {
	task := &models.Task{}

	updateRes := r.db.Model(task).Where("id = ?", taskId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", taskId).First(task)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return task, nil
}

func (r *TaskRepository) DeleteOne(ctx context.Context, taskId uuid.UUID) error {
	res := r.db.Delete(&models.Task{}, taskId)

	return res.Error
}


func NewTaskRepository(db *gorm.DB) models.TaskRepository {
	return &TaskRepository{
		db: db,
	}
}