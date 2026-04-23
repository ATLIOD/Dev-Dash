package repositories

import (
	"DevDash/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository interface {
	GetByID(ctx context.Context, id string) (*models.Project, error)
	Create(ctx context.Context, project *models.Project) error
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id string) error
}

type projectRepository struct {
	db *pgxpool.Pool
}

func (r *projectRepository) GetByID(ctx context.Context, id string) (*models.Project, error) {
	query := `
		SELECT id, name, description, status, stack, repository_url, deployment_url, user_id, created_at, updated_at
		FROM projects
		WHERE id = $1
	`
	var project models.Project
	err := r.db.QueryRow(ctx, query, id).Scan(&project.ID, &project.Name, &project.Description, &project.Status, &project.Stack, &project.RepositoryURL, &project.DeploymentURL, &project.UserID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) error {
	query := `
		INSERT INTO projects (id, name, description, status, stack, repository_url, deployment_url, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, project.Name, project.Description, project.Status, project.Stack, project.RepositoryURL, project.DeploymentURL, project.UserID).Scan(&project.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *projectRepository) Update(ctx context.Context, user *models.Project) error {
	query := `
		UPDATE projects
		SET name = $1, description = $2, status = $3, stack = $4, repository_url = $5, deployment_url = $6, updated_at = NOW()
		WHERE id = $8
		`
	_, err := r.db.Exec(ctx, query, user.Name, user.Description, user.Status, user.Stack, user.RepositoryURL, user.DeploymentURL, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *projectRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM projects
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
