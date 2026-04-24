package models

import "time"

type Project struct {
	ID            int64     `json:"id"`
	UUID          string    `json:"uuid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Status        string    `json:"status"` // e.g., "active", "completed", "archived"
	Stack         string    `json:"stack"`
	RepositoryURL string    `json:"repository_url"`
	DeploymentURL string    `json:"deployment_url"`
	UserID        int64     `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateProjectRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	Stack         string `json:"stack"`
	RepositoryURL string `json:"repository_url"`
	DeploymentURL string `json:"deployment_url"`
	UserID        int64  `json:"user_id"`
}

type UpdateProjectRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	Stack         string `json:"stack"`
	RepositoryURL string `json:"repository_url"`
	DeploymentURL string `json:"deployment_url"`
}

type ProjectResponse struct {
	ID            string    `json:"uuid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	Stack         string    `json:"stack"`
	RepositoryURL string    `json:"repository_url"`
	DeploymentURL string    `json:"deployment_url"`
	UserID        int64     `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (p *Project) ToResponse() ProjectResponse {
	return ProjectResponse{
		ID:            p.UUID,
		Name:          p.Name,
		Description:   p.Description,
		Status:        p.Status,
		Stack:         p.Stack,
		RepositoryURL: p.RepositoryURL,
		DeploymentURL: p.DeploymentURL,
		UserID:        p.UserID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
