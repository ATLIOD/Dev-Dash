package services

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
)

type ProjectService interface {
	GetByID(ctx context.Context, id string) (*models.ProjectResponse, error)
	Create(ctx context.Context, req models.CreateProjectRequest) (*models.ProjectResponse, error)
	Update(ctx context.Context, id string, req models.UpdateProjectRequest) (*models.ProjectResponse, error)
	Delete(ctx context.Context, id string) error
}

type projectService struct {
	projectRepo repositories.ProjectRepository
}

func (s *projectService) GetByID(ctx context.Context, id string) (*models.ProjectResponse, error) {
}

func (s *projectService) Create(ctx context.Context, req models.CreateProjectRequest) (*models.ProjectResponse, error) {
}

func (s *projectService) Update(ctx context.Context, id string, req models.UpdateProjectRequest) (*models.ProjectResponse, error) {
}

func (s *projectService) Delete(ctx context.Context, id string) error {
}
