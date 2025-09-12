package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meshackkiplimo/Taski/database"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Create Fiber app
	app := fiber.New()

	// Basic root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Task Manager API")
	})

	// Start server
	app.Listen(":5000")
}
