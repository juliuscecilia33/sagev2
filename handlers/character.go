package handlers

import (
	"context"
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
	return nil
}

func (h *CharacterHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func NewCharacterHandler(router fiber.Router, repository models.CharacterRepository) {
	handler := &CharacterHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:characterId", handler.GetOne)
}