package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/config"
	"github.com/juliuscecilia33/sagev2/db"
	"github.com/juliuscecilia33/sagev2/handlers"
	"github.com/juliuscecilia33/sagev2/middlewares"
	"github.com/juliuscecilia33/sagev2/repositories"
	"github.com/juliuscecilia33/sagev2/services"
)

func main() {
	envConfig := config.NewEnvConfig() // Load config from environment variables
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "Sage",
		ServerHeader: "Fiber",
	})

	// Repositories
	characterRepository := repositories.NewCharacterRepository(db)
	itemRepository := repositories.NewItemRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	userKidRepository := repositories.NewUserKidRepository(db)
	quizRepository := repositories.NewQuizRepository(db)
	rewardRepository := repositories.NewRewardRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewCharacterHandler(privateRoutes.Group("/character"), characterRepository)
	handlers.NewItemHandler(privateRoutes.Group("/item"), itemRepository)
	handlers.NewUserKidHandler(privateRoutes.Group("/userkid"), userKidRepository)
	handlers.NewQuizHandler(privateRoutes.Group("/quiz"), quizRepository)
	handlers.NewRewardHandler(privateRoutes.Group("/reward"), rewardRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}