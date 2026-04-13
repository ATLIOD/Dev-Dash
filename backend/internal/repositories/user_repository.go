package repositories

import (
	"DevDash/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	List(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

// these should just be doing database stuff and passing results back
//  from my understanding: no business logic here
func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}
func (r *userRepository) List(ctx context.Context) ([]models.User, error)     { return nil, nil }
func (r *userRepository) Create(ctx context.Context, user *models.User) error { return nil }
func (r *userRepository) Update(ctx context.Context, user *models.User) error { return nil }
func (r *userRepository) Delete(ctx context.Context, id string) error         { return nil }
