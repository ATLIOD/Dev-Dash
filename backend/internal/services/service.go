package services

import "DevDash/internal/repositories"

type Service struct {
	User UserService
}

func New(repo *repositories.Repository) *Service {
	return &Service{
		User: &userService{userRepo: repo.User},
	}
}
