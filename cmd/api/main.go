package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/config"
	"github.com/juliuscecilia33/sagev2/db"
	"github.com/juliuscecilia33/sagev2/handlers"
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

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middleware.AuthProtected(db))

	// Handlers
	handlers.NewCharacterHandler(server.Group("/character"), characterRepository)
	handlers.NewItemHandler(server.Group("/item"), itemRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}