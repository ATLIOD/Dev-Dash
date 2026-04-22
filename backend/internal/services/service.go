package services

import "DevDash/internal/repositories"

type Service struct {
	User    UserService
	Project ProjectService
}

func New(repo *repositories.Repository) *Service {
	return &Service{
		User:    &userService{userRepo: repo.User},
		Project: &projectService{projectRepo: repo.Project},
	}
}
