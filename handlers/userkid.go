package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/models"
	"github.com/google/uuid"
	"time"
)

type UserKidHandler struct {
	repository models.UserKidRepository
}

// getContext returns a context with a timeout
func (h *UserKidHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *UserKidHandler) GetAllParentKids(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	parentId := ctx.Params("parentId")
	// Parse the UUID from the string
	parsedParentID, err := uuid.Parse(parentId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid parent ID",
		})
	}

	kids, err := h.repository.GetAllParentKids(ctxWithTimeout, parsedParentID)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all kids for parent",
		"data":    kids,
	})
}

func (h *UserKidHandler) GetOne(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	kidId := ctx.Params("userKidId")
	// Parse the UUID from the string
	parsedKidID, err := uuid.Parse(kidId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user kid ID",
		})
	}

	kid, err := h.repository.GetOne(ctxWithTimeout, parsedKidID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved user kid",
		"data":    kid,
	})
}

func (h *UserKidHandler) CreateOne(ctx *fiber.Ctx) error {
	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	kid := &models.UserKid{}
	if err := ctx.BodyParser(kid); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	createdKid, err := h.repository.CreateOne(ctxWithTimeout, kid)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "user kid created",
		"data":    createdKid,
	})
}

func (h *UserKidHandler) UpdateOne(ctx *fiber.Ctx) error {
	kidId := ctx.Params("userKidId")
	// Parse the UUID from the string
	parsedKidID, err := uuid.Parse(kidId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user kid ID",
		})
	}

	updateData := make(map[string]interface{})
	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	updatedKid, err := h.repository.UpdateOne(ctxWithTimeout, parsedKidID, updateData)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "user kid updated",
		"data":    updatedKid,
	})
}

func (h *UserKidHandler) DeleteOne(ctx *fiber.Ctx) error {
	kidId := ctx.Params("userKidId")
	// Parse the UUID from the string
	parsedKidID, err := uuid.Parse(kidId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user kid ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedKidID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewUserKidHandler(router fiber.Router, repository models.UserKidRepository) {
	handler := &UserKidHandler{repository: repository}

	router.Get("/parentkids/:parentId", handler.GetAllParentKids)
	router.Post("/", handler.CreateOne)
	router.Get("/:userKidId", handler.GetOne)
	router.Put("/:userKidId", handler.UpdateOne)
	router.Delete("/:userKidId", handler.DeleteOne)
}
