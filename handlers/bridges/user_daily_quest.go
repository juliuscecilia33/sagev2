package bridges

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges" // import models/bridges for UserQuiz
)

type UserDailyQuestHandler struct {
	repository bridges.UserDailyQuestRepository
}

func (h *UserDailyQuestHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

