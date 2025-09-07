package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/task-manager/database"
    "github.com/yourusername/task-manager/handlers"
)

func main() {
    // Initialize database
    database.Connect()

    // Create Fiber app
    app := fiber.New()

    // Define routes
    app.Post("/tasks", handlers.CreateTask)
    app.Get("/tasks", handlers.GetTasks)
    app.Get("/tasks/:id", handlers.GetTask)
    app.Put("/tasks/:id", handlers.UpdateTask)
    app.Delete("/tasks/:id", handlers.DeleteTask)

    // Start server
    app.Listen(":3000")
}