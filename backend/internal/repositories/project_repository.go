package repositories

import (
	"DevDash/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository interface {
	GetByID(ctx context.Context, id string) (*models.Project, error)
	GetByEmail(ctx context.Context, email string) (*models.Project, error)
	Create(ctx context.Context, project *models.Project) error
	Update(ctx context.Context, project *models.Project) error
	Delete(ctx context.Context, id string) error
}

type projectRepository struct {
	db *pgxpool.Pool
}

func (r *projectRepository) GetByID(ctx context.Context, id string) (*models.Project, error) {
}

func (r *projectRepository) GetByEmail(ctx context.Context, email string) (*models.Project, error) {
}

func (r *projectRepository) Create(ctx context.Context, user *models.Project) error {
}

func (r *projectRepository) Update(ctx context.Context, user *models.Project) error {
}

func (r *projectRepository) Delete(ctx context.Context, id string) error {
}
