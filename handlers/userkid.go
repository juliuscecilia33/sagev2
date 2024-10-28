package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/models"
)

type UserKidHandler struct {
	repository models.UserKidRepository
}

func (h *UserKidHandler) GetAllParentKids(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	parentId, _ := strconv.Atoi(ctx.Params("parentId"))

	kids, err := h.repository.GetAllParentKids(context, uint(parentId))

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK). JSON(&fiber.Map{
		"status": "success",
		"message": "retrieved all kids for parent",
		"data": kids,
	})
}

func (h *UserKidHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	kidId, _ := strconv.Atoi(ctx.Params("userKidId"))

	kid, err := h.repository.GetOne(context, uint(kidId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "retrieved user kid",
		"data": kid,
	})
}

func (h *UserKidHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	kid := &models.UserKid{}

	if err := ctx.BodyParser(kid); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
			"data": nil,
		})
	}


	kid, err := h.repository.CreateOne(context, kid)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "user kid created",
		"data": kid,
	})
}

func (h *UserKidHandler ) UpdateOne(ctx *fiber.Ctx) error {
	kidId, _ := strconv.Atoi(ctx.Params("userKidId"))

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
	
	kid, err := h.repository.UpdateOne(context, uint(kidId), updateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}


	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "user kid updated",
		"data": kid,
	})

}

func (h* UserKidHandler) DeleteOne(ctx *fiber.Ctx) error {
	kidId, _ := strconv.Atoi(ctx.Params("userKidId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context, uint(kidId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// TODO: Validate kid? check items validation in valliant tutorial

func NewUserKidHandler(router fiber.Router, repository models.UserKidRepository) {
	handler := &UserKidHandler{
		repository: repository,
	}

	router.Get("/:parentId", handler.GetAllParentKids)
	router.Post("/", handler.CreateOne)
	router.Get("/:userKidId", handler.GetOne)
	router.Put("/:userKidId", handler.UpdateOne)
	router.Delete("/:userKidId", handler.DeleteOne)
}