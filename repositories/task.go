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