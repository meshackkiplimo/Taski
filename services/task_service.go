package services

import (
	"errors"
	"github.com/Taski/models"
	"github.com/Taski/repository"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(req CreateTaskRequest) (*models.Task, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	if task.Status == "" {
		task.Status = "pending"
	}

	err := s.taskRepo.Create(task)
	if err != nil {
		return nil, errors.New("could not create task")
	}

	return task, nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.taskRepo.FindAll()
}

func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) UpdateTask(id string, req UpdateTaskRequest) (*models.Task, error) {
	task, err := s.taskRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		task.Status = req.Status
	}

	err = s.taskRepo.Update(task)
	if err != nil {
		return nil, errors.New("could not update task")
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	task, err := s.taskRepo.FindByID(id)
	if err != nil {
		return errors.New("task not found")
	}

	err = s.taskRepo.Delete(task)
	if err != nil {
		return errors.New("could not delete task")
	}

	return nil
}
