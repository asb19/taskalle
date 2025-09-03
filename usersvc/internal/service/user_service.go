package service

import (
	"context"

	"github.com/asb19/usersvc/internal/model"
	"github.com/asb19/usersvc/internal/repo"
	"github.com/google/uuid"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.UserPublicInfo, error)
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.UserPublicInfo, error) {
	return s.repo.GetByID(ctx, id)
}

// Define domain-specific errors
var ErrInvalidTask = &ServiceError{"invalid task data"}

type ServiceError struct {
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}
