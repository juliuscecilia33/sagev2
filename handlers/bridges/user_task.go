package bridges

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges"
)


type UserTaskHandler struct {
	repository bridges.UserTaskRepository
}

func (h *UserTaskHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *UserTaskHandler) GetAllByUser(ctx *fiber.Ctx) error {
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

	specific_user_tasks, err := h.repository.GetAllByUser(ctxWithTimeout, parsedUserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all tasks for specific user",
		"data":    specific_user_tasks,
	})
}

func (h *UserTaskHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_tasks, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all user tasks",
		"data":    user_tasks,
	})
}

func (h *UserTaskHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userTaskId, err := uuid.Parse(ctx.Params("userTaskId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user Character ID",
		})
	}

	userTask, err := h.repository.GetOne(context, userTaskId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved user task",
		"data":    userTask,
	})
}

func (h *UserTaskHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_task := &bridges.UserTask{}

	if err := ctx.BodyParser(user_task); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user_task, err := h.repository.CreateOne(context, user_task)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created user task",
		"data":    user_task,
	})
}

func (h *UserTaskHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userTaskId, err := uuid.Parse(ctx.Params("userTaskId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user task ID",
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

	user_task, err := h.repository.UpdateOne(context, userTaskId, updateData) 

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "user task updated",
		"data":    user_task,
	})
}

func (h *UserTaskHandler) DeleteOne(ctx *fiber.Ctx) error {
	userTaskId := ctx.Params("userTaskId")
	// Parse the UUID from the string
	parsedUserTaskID, err := uuid.Parse(userTaskId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user task ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedUserTaskID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewUserTaskHandler(router fiber.Router, repository bridges.UserTaskRepository) {
	handler := &UserTaskHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:userTaskId", handler.GetOne)
	router.Get("/user/:userId", handler.GetAllByUser)
	router.Put("/:userTaskId", handler.UpdateOne)
	router.Delete("/:userTaskId", handler.DeleteOne)
}