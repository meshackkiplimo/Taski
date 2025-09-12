package controllers

import (
	"github.com/Taski/services"
	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(taskService *services.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	var req services.CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	task, err := tc.taskService.CreateTask(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (tc *TaskController) GetTasks(c *fiber.Ctx) error {
	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(tasks)
}

func (tc *TaskController) GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := tc.taskService.GetTaskByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(task)
}

func (tc *TaskController) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var req services.UpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	task, err := tc.taskService.UpdateTask(id, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(task)
}

func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	err := tc.taskService.DeleteTask(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
