package services

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
)

type UserService interface {
	GetByID(ctx context.Context, id string) (*models.UserResponse, error)
	List(ctx context.Context) ([]models.UserResponse, error)
	Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.UserResponse, error) {
	return nil, nil
}
func (s *userService) List(ctx context.Context) ([]models.UserResponse, error) { return nil, nil }
func (s *userService) Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	return nil, nil
}

func (s *userService) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.UserResponse, error) {
	return nil, nil
}
func (s *userService) Delete(ctx context.Context, id string) error { return nil }
