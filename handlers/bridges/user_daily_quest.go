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

func (h *UserDailyQuestHandler) GetAllByUser(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	userId := ctx.Params("userId")
	// Parse the UUID from the string
	parsedUserID, err := uuid.Parse(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user ID",
		})
	}

	specific_user_daily_quests, err := h.repository.GetAllByUser(ctxWithTimeout, parsedUserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all daily quests for specific user",
		"data":    specific_user_daily_quests,
	})
}

func (h *UserDailyQuestHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_daily_quests, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all user daily quests",
		"data":    user_daily_quests,
	})
}

func (h *UserDailyQuestHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userDailyQuestId, err := uuid.Parse(ctx.Params("userDailyQuestId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user daily quest ID",
		})
	}

	userDailyQuest, err := h.repository.GetOne(context, userDailyQuestId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved user daily quest",
		"data":    userDailyQuest,
	})
}

func (h *UserDailyQuestHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_daily_quest := &bridges.UserDailyQuest{}

	if err := ctx.BodyParser(user_daily_quest); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user_daily_quest, err := h.repository.CreateOne(context, user_daily_quest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created user daily quest",
		"data":    user_daily_quest,
	})
}

func (h *UserDailyQuestHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userQuizId, err := uuid.Parse(ctx.Params("userQuizId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user daily quest ID",
		})
	}

	updateData := make(map[string]interface{})

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user_daily_quest, err := h.repository.UpdateOne(context, userQuizId, updateData) 
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "user daily quest updated",
		"data":    user_daily_quest,
	})
}

func (h *UserDailyQuestHandler) DeleteOne(ctx *fiber.Ctx) error {
	userDailyQuestId := ctx.Params("userQuizId")
	// Parse the UUID from the string
	parsedUserDailyQuestID, err := uuid.Parse(userDailyQuestId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user daily quest ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedUserDailyQuestID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewUserDailyQuestHandler(router fiber.Router, repository bridges.UserDailyQuestRepository) {
	handler := &UserDailyQuestHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:userDailyQuestId", handler.GetOne)
	router.Get("/user/:userId", handler.GetAllByUser)
	router.Put("/:userDailyQuestId", handler.UpdateOne)
	router.Delete("/:userDailyQuestId", handler.DeleteOne)
}
