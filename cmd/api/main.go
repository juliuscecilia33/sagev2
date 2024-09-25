package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliuscecilia33/sagev2/config"
	"github.com/juliuscecilia33/sagev2/db"
	"github.com/juliuscecilia33/sagev2/handlers"
	"github.com/juliuscecilia33/sagev2/repositories"
)

func main() {
	envConfig := config.NewEnvConfig() // Load config from environment variables
	db := db.Init(envConfig)

	app := fiber.New(fiber.Config{
		AppName: "Sage",
		ServerHeader: "Fiber",
	})

	// Repositories
	characterRepository := repositories.NewCharacterRepository(nil)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewCharacterHandler(server.Group("/character"), characterRepository)

	app.Listen(":3000")
}