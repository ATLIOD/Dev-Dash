package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	User    UserRepository
	Project ProjectRepository
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		User:    &userRepository{db: db},
		Project: &projectRepository{db: db},
	}
}
