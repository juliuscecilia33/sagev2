package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models"
)

type DailyQuestHandler struct {
	repository models.DailyQuestRepository
}

func (h *DailyQuestHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *DailyQuestHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	dailyQuests, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all daily quests",
		"data":    dailyQuests,
	})
}

func (h *DailyQuestHandler) GetByDate(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	questDate := ctx.Params("questDate")
	if questDate == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "quest date is required",
		})
	}

	// Fetch quests by date from the repository
	dailyQuests, err := h.repository.GetByDate(context, questDate)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved daily quests for date",
		"data":    dailyQuests,
	})
}

func (h *DailyQuestHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	dailyQuest := &models.DailyQuest{}

	if err := ctx.BodyParser(dailyQuest); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	dailyQuest, err := h.repository.CreateOne(context, dailyQuest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created daily quest",
		"data":    dailyQuest,
	})
}

func (h *DailyQuestHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	dailyQuestId, err := uuid.Parse(ctx.Params("dailyQuestId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid daily quest ID",
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

	dailyQuest, err := h.repository.UpdateOne(context, dailyQuestId, updateData) 
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "daily quest updated",
		"data":    dailyQuest,
	})
}

func (h *DailyQuestHandler) DeleteOne(ctx *fiber.Ctx) error {
	dailyQuestId := ctx.Params("dailyQuestId")
	// Parse the UUID from the string
	parsedDailyQuestID, err := uuid.Parse(dailyQuestId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid daily quest ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedDailyQuestID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *DailyQuestHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	dailyQuestId, err := uuid.Parse(ctx.Params("dailyQuestId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid daily quest ID",
		})
	}

	dailyQuest, err := h.repository.GetOne(context, dailyQuestId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved daily quest",
		"data":    dailyQuest,
	})
}

func NewDailyQuestHandler(router fiber.Router, repository models.DailyQuestRepository) {
	handler := &DailyQuestHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:dailyQuestId", handler.GetOne)
	router.Get("/date/:questDate", handler.GetByDate)
	router.Put("/:dailyQuestId", handler.UpdateOne)
	router.Delete("/:dailyQuestId", handler.DeleteOne)
}
