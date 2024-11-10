package bridges

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/juliuscecilia33/sagev2/models/bridges" // import models/bridges for UserQuiz
)

type UserQuizHandler struct {
	repository bridges.UserQuizRepository
}

func (h *UserQuizHandler) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (h *UserQuizHandler) GetAllByUser(ctx *fiber.Ctx) error {
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

	specific_user_quizzes, err := h.repository.GetAllByUser(ctxWithTimeout, parsedUserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all quizzes for specific user",
		"data":    specific_user_quizzes,
	})
}

func (h *UserQuizHandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_quizzes, err := h.repository.GetMany(context)

	// If there is an error
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved all user quizzes",
		"data":    user_quizzes,
	})
}

func (h *UserQuizHandler) GetOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userQuizId, err := uuid.Parse(ctx.Params("userQuizId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user Quiz ID",
		})
	}

	userQuiz, err := h.repository.GetOne(context, userQuizId) // Pass UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "retrieved user quiz",
		"data":    userQuiz,
	})
}

func (h *UserQuizHandler) CreateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	user_quiz := &bridges.UserQuiz{}

	if err := ctx.BodyParser(user_quiz); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user_quiz, err := h.repository.CreateOne(context, user_quiz)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "created user quiz",
		"data":    user_quiz,
	})
}

func (h *UserQuizHandler) UpdateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	userQuizId, err := uuid.Parse(ctx.Params("userQuizId")) // Change to parse UUID
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user quiz ID",
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

	user_quiz, err := h.repository.UpdateOne(context, userQuizId, updateData) 
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "user quiz updated",
		"data":    user_quiz,
	})
}

func (h *UserQuizHandler) DeleteOne(ctx *fiber.Ctx) error {
	userQuizId := ctx.Params("userQuizId")
	// Parse the UUID from the string
	parsedUserQuizID, err := uuid.Parse(userQuizId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "invalid user quiz ID",
		})
	}

	ctxWithTimeout, cancel := h.getContext()
	defer cancel()

	if err := h.repository.DeleteOne(ctxWithTimeout, parsedUserQuizID); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewUserQuizHandler(router fiber.Router, repository bridges.UserQuizRepository) {
	handler := &UserQuizHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:userQuizId", handler.GetOne)
	router.Get("/user/:userId", handler.GetAllByUser)
	router.Put("/:userQuizId", handler.UpdateOne)
	router.Delete("/:userQuizId", handler.DeleteOne)
}
