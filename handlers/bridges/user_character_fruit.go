package bridges

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
)

type UserCharacterFruitHandler struct {
	repository bridges.UserCharacterFruitRepository
}

func (h *UserCharacterFruitHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *UserCharacterFruitHandler) GetAllByUser(ctx *fiber.Ctx) error {
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

	specific_user_character_fruits, err := h.repository.GetAllByUser(ctxWithTimeout, parsedUserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all fruits for character for specific user",
		"data":    specific_user_character_fruits,
	})
}

func (h *UserCharacterFruitHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	specific_user_character_fruits, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all fruits for user characters",
		"data":    specific_user_character_fruits,
	})
}

func (h *UserCharacterFruitHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userCharacterFruitId, err := uuid.Parse(ctx.Params("userCharacterFruitId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user Character ID",
		})
	}

	userCharacterFruit, err := h.repository.GetOne(context, userCharacterFruitId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved specific fruit for user character",
		"data":    userCharacterFruit,
	})
}

func (h *UserCharacterFruitHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_character_fruit := &bridges.UserCharacterFruit{}

	if err := ctx.BodyParser(user_character_fruit); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user_character_fruit, err := h.repository.CreateOne(context, user_character_fruit)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created fruit for user character",
		"data":    user_character_fruit,
	})
}

func (h *UserCharacterFruitHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userCharacterFruitId, err := uuid.Parse(ctx.Params("userCharacterFruitId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user character fruit ID",
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

	user_character_fruit, err := h.repository.UpdateOne(context, userCharacterFruitId, updateData) 
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "fruit for user character updated",
		"data":    user_character_fruit,
	})
}

func (h *UserCharacterFruitHandler) DeleteOne(ctx *fiber.Ctx) error {
	userCharacterFruitId := ctx.Params("userCharacterFruitId")
	// Parse the UUID from the string
	parsedUserCharacterFruitID, err := uuid.Parse(userCharacterFruitId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user character ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedUserCharacterFruitID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewUserCharacterFruitHandler(router fiber.Router, repository bridges.UserCharacterFruitRepository) {
	handler := &UserCharacterFruitHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:userCharacterFruitId", handler.GetOne)
	router.Get("/user/:userId", handler.GetAllByUser)
	router.Put("/:userCharacterFruitId", handler.UpdateOne)
	router.Delete("/:userCharacterFruitId", handler.DeleteOne)
}