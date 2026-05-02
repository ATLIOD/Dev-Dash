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

func (s *projectService) GetByUUID(ctx context.Context, id string) (*models.ProjectResponse, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := project.ToResponse()
	return &resp, nil
}

func (s *projectService) Create(ctx context.Context, req models.CreateProjectRequest) (*models.ProjectResponse, error) {
	project := &models.Project{
		Name:          req.Name,
		Description:   req.Description,
		Status:        req.Status,
		Stack:         req.Stack,
		RepositoryURL: req.RepositoryURL,
		DeploymentURL: req.DeploymentURL,
		UserID:        req.UserID,
	}
	err := s.projectRepo.Create(ctx, project)
	if err != nil {
		return nil, err
	}
	resp := project.ToResponse()
	return &resp, nil
}

func (s *projectService) Update(ctx context.Context, id string, req models.UpdateProjectRequest) (*models.ProjectResponse, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	project.Name = req.Name
	project.Description = req.Description
	project.Status = req.Status
	project.Stack = req.Stack
	project.RepositoryURL = req.RepositoryURL
	project.DeploymentURL = req.DeploymentURL

	err = s.projectRepo.Update(ctx, project)
	if err != nil {
		return nil, err
	}
	resp := project.ToResponse()
	return &resp, nil
}

func (s *projectService) Delete(ctx context.Context, id string) error {
	err := s.projectRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
