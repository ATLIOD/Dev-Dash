package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User UserRepository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: &userRepository{db: db},
	}
}
