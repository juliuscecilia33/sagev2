package bridges

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
	"gorm.io/gorm"
)

type UserTaskRepository struct {
	db *gorm.DB
}

func (r *UserTaskRepository) GetAllByUser(ctx context.Context, userId uuid.UUID) ([]*bridges.UserTask, error) {
	specific_user_tasks := []*bridges.UserTask{}

	res := r.db.Model(&bridges.UserTask{}).Where("user_id = ?", userId).Preload("User").Preload("Task").Preload("Task.Reward").Find(&specific_user_tasks)

	if res.Error != nil {
		return nil, res.Error
	}

	return specific_user_tasks, nil
}

func (r *UserTaskRepository) GetMany(ctx context.Context) ([]*bridges.UserTask, error) {
	user_tasks := []*bridges.UserTask{}

	res := r.db.Model(&bridges.UserTask{}).Preload("User").Preload("Task").Preload("Task.Reward").Order("updated_at desc").Find(&user_tasks)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_tasks, nil
}

func (r *UserTaskRepository) GetOne(ctx context.Context, userTaskId uuid.UUID) (*bridges.UserTask, error) {
	user_task := &bridges.UserTask{}

	res := r.db.Model(user_task).Preload("User").Preload("Task").Preload("Task.Reward").Where("id = ?", userTaskId).First(user_task)

	if res.Error != nil {
		return nil, res.Error
	}

	return user_task, nil
}

func (r *UserTaskRepository) CreateOne(ctx context.Context, user_task *bridges.UserTask) (*bridges.UserTask, error) {
	res := r.db.Create(user_task)

	if res.Error != nil {
		return nil, res.Error
	}

	if err := r.db.Preload("User").Preload("Task").First(user_task, user_task.ID).Error; err != nil {
		return nil, err
	}

	return user_task, nil
}

func (r UserTaskRepository) UpdateOne(ctx context.Context, userTaskId uuid.UUID, updateData map[string]interface{}) (*bridges.UserTask, error) {
	user_task := &bridges.UserTask{}

	updateRes := r.db.Model(user_task).Where("id = ?", userTaskId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getRes := r.db.Where("id = ?", userTaskId).First(user_task)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return user_task, nil
}


func (r *UserTaskRepository) DeleteOne(ctx context.Context, userTaskId uuid.UUID) error {
	res := r.db.Delete(&bridges.UserTask{}, userTaskId)

	return res.Error
}

func NewUserTaskRepository(db *gorm.DB) bridges.UserTaskRepository {
	return &UserTaskRepository{
		db: db,
	}
}