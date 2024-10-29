package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid" // Make sure to import the uuid package
	"github.com/juliuscecilia33/sagev2/models"
)

type ItemHandler struct {
	repository models.ItemRepository
}

func (h *ItemHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	items, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all items",
		"data":    items,
	})
}

func (h *ItemHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	itemId, err := uuid.Parse(ctx.Params("itemId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid item ID",
		})
	}

	item, err := h.repository.GetOne(context, itemId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved item",
		"data":    item,
	})
}

func (h *ItemHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	item := &models.Item{}

	if err := ctx.BodyParser(item); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	item, err := h.repository.CreateOne(context, item)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created item",
		"data":    item,
	})
}

func (h *ItemHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	itemId, err := uuid.Parse(ctx.Params("itemId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid item ID",
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

	item, err := h.repository.UpdateOne(context, itemId, updateData) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "item updated",
		"data":    item,
	})
}

func NewItemHandler(router fiber.Router, repository models.ItemRepository) {
	handler := &ItemHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:itemId", handler.GetOne)
	router.Put("/:itemId", handler.UpdateOne)
}
