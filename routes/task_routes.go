package routes

import (
	"github.com/Taski/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App, taskController *controllers.TaskController) {
	api := app.Group("/api/v1")

	api.Post("/tasks", taskController.CreateTask)
	api.Get("/tasks", taskController.GetTasks)
	api.Get("/tasks/:id", taskController.GetTask)
	api.Put("/tasks/:id", taskController.UpdateTask)
	api.Delete("/tasks/:id", taskController.DeleteTask)
}
