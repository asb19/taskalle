package service

import (
	"context"

	"github.com/asb19/tasksvc/internal/model"
	"github.com/asb19/tasksvc/internal/repo"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *model.Task) (model.Task, error)
	GetTasks(ctx context.Context, status string, page, limit int) ([]model.Task, error)
	GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type taskService struct {
	repo repo.TaskRepository
}

func NewTaskService(repo repo.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) (model.Task, error) {
	if task.Title == "" {
		return model.Task{}, ErrInvalidTask
	}
	return s.repo.Create(ctx, task)
}

func (s *taskService) GetTasks(ctx context.Context, status string, page, limit int) ([]model.Task, error) {
	return s.repo.GetAll(ctx, status, page, limit)
}

func (s *taskService) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *taskService) UpdateTask(ctx context.Context, task *model.Task) error {
	return s.repo.Update(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// Define domain-specific errors
var ErrInvalidTask = &ServiceError{"invalid task data"}

type ServiceError struct {
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}
