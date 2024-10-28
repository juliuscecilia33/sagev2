package handlers

import (
	"context"
	"strconv"
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

	items, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK). JSON(&fiber.Map{
		"status": "success",
		"message": "retrieved all items",
		"data": items,
	})
}

func (h *ItemHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	itemId, _ := strconv.Atoi(ctx.Params("itemId"))

	item, err := h.repository.GetOne(context, uint(itemId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "retrieved item",
		"data": item,
	})
}

func (h *ItemHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	item := &models.Item{}

	if err := ctx.BodyParser(item); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
			"data": nil,
		})
	}


	item, err := h.repository.CreateOne(context, item)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "created item",
		"data": item,
	})
}

func (h *ItemHandler ) UpdateOne(ctx *fiber.Ctx) error {
	itemId, _ := strconv.Atoi(ctx.Params("itemId"))

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
	
	item, err := h.repository.UpdateOne(context, uint(itemId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}


	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "item updated",
		"data": item,
	})

}

// ValidateOne - 1:29:00

func NewItemHandler(router fiber.Router, repository models.ItemRepository) {
	handler := &ItemHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:itemId", handler.GetOne)
	router.Put("/:itemId", handler.UpdateOne)
} 