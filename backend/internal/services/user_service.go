package services

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
)

type UserService interface {
	GetByID(ctx context.Context, id string) (*models.UserResponse, error)
	Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

// all of these basically just call the repo and do the other logic that is needed if any
// they do not ever touch the database directly
func (s *userService) GetByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := user.ToResponse()
	return &resp, nil
}
func (s *userService) Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := user.ToResponse()
	return &resp, nil
}

func (s *userService) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Name = req.Name
	user.Email = req.Email
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := user.ToResponse()
	return &resp, nil
}
func (s *userService) Delete(ctx context.Context, id string) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
