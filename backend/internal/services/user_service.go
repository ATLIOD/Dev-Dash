package services

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"DevDash/pkg/utils"
	"context"
	"errors"
	"strings"
)

type UserService interface {
	GetByID(ctx context.Context, id string) (*models.UserResponse, error)
	GetByEmail(ctx context.Context, email string) (*models.UserResponse, error)
	Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.UserResponse, error)
	Delete(ctx context.Context, id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := user.ToResponse()
	return &resp, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	resp := user.ToResponse()
	return &resp, nil
}

func (s *userService) Create(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
	password := strings.TrimSpace(req.Password)
	if password == "" {
		return nil, errors.New("password is required")
	}

	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}
	err = s.userRepo.Create(ctx, user)
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
