package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/models"
)

type CharacterHandler struct {
	repository models.CharacterRepository
}

func (h *CharacterHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	characters, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK). JSON(&fiber.Map{
		"status": "success",
		"message": "",
		"data": characters,
	})
}

func (h *CharacterHandler) GetOne(ctx *fiber.Ctx) error {
	characterId, _ := strconv.Atoi(ctx.Params("characterId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	character, err := h.repository.GetOne(context, uint(characterId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "",
		"data": character,
	})
}

func (h *CharacterHandler) CreateOne(ctx *fiber.Ctx) error {
	character := &models.Character{}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(character); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
			"data": nil,
		})
	}

	character, err := h.repository.CreateOne(context, character)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}


	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "Character created",
		"data": character,
	})
}

func (h *CharacterHandler) UpdateOne(ctx *fiber.Ctx) error {
	characterId, _ := strconv.Atoi(ctx.Params("characterId"))

	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
			"data": nil,
		})
	}
	
	character, err := h.repository.UpdateOne(context, uint(characterId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}


	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "Character created",
		"data": character,
	})

}

func (h* CharacterHandler) DeleteOne(ctx *fiber.Ctx) error {
	characterId, _ := strconv.Atoi(ctx.Params("characterId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context, uint(characterId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewCharacterHandler(router fiber.Router, repository models.CharacterRepository) {
	handler := &CharacterHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:characterId", handler.GetOne)
	router.Put("/:characterId", handler.UpdateOne)
	router.Delete("/:characterId", handler.DeleteOne)
}