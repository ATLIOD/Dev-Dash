package repositories

import (
	"DevDash/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByUUID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, uuid, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUUID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, uuid, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE uuid = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, uuid, name, email, password_hash, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, uuid
	`
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.PasswordHash).Scan(&user.ID, &user.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2
		WHERE uuid = $3
	`
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE uuid = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
