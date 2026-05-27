package utils

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
)

func Setup() *repositories.Repository {
	db := models.NewMockDB()
	repo := repositories.NewMockRepo(db)

	return repo
}
