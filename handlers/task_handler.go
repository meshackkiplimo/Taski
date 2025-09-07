package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/yourusername/task-manager/database"
    "github.com/yourusername/task-manager/models"
)

// CreateTask creates a new task
func CreateTask(c *fiber.Ctx) error {
    task := new(models.Task)
    if err := c.BodyParser(task); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if task.Title == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
    }

    if err := database.DB.Create(&task).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create task"})
    }

    return c.Status(fiber.StatusCreated).JSON(task)
}

// GetTasks retrieves all tasks
func GetTasks(c *fiber.Ctx) error {
    var tasks []models.Task
    if err := database.DB.Find(&tasks).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve tasks"})
    }
    return c.JSON(tasks)
}

// GetTask retrieves a task by ID
func GetTask(c *fiber.Ctx) error {
    id := c.Params("id")
    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
    }
    return c.JSON(task)
}

// UpdateTask updates a task by ID
func UpdateTask(c *fiber.Ctx) error {
    id := c.Params("id")
    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
    }

    updateData := new(models.Task)
    if err := c.BodyParser(updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    // Update fields if provided
    if updateData.Title != "" {
        task.Title = updateData.Title
    }
    if updateData.Description != "" {
        task.Description = updateData.Description
    }
    if updateData.Status != "" {
        task.Status = updateData.Status
    }

    if err := database.DB.Save(&task).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not update task"})
    }

    return c.JSON(task)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *fiber.Ctx) error {
    id := c.Params("id")
    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
    }

    if err := database.DB.Delete(&task).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete task"})
    }

    return c.SendStatus(fiber.StatusNoContent)
}