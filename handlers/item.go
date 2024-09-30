package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/models"
)

type ItemHandler struct {
	repository models.ItemRepository
}

func (h *ItemHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
}

func NewItemHandler(router fiber.Router, repository models.ItemRepository) {
	handler := &ItemHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:itemId", handler.GetOne)
} 