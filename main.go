package main

import (
	"github.com/Taski/controllers"
	"github.com/Taski/database"
	"github.com/Taski/repository"
	"github.com/Taski/routes"
	"github.com/Taski/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Initialize layers
	taskRepo := repository.NewTaskRepository(database.DB)
	taskService := services.NewTaskService(taskRepo)
	taskController := controllers.NewTaskController(taskService)

	// Create Fiber app
	app := fiber.New()

	// Basic root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Task Manager API")
	})

	// Setup routes
	routes.SetupTaskRoutes(app, taskController)

	// Start server
	app.Listen(":5000")
}
