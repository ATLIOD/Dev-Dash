package utils

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
	"log"
)

func UserSetup(repo *repositories.Repository, test string) *models.User {
	if test == "" {
		test = "default"
	}
	user := &models.User{
		Name:         "test-user-" + test,
		Email:        "test-user-" + test + "@example.com",
		PasswordHash: "password123",
	}

	err := repo.User.Create(context.Background(), user)
	if err != nil {
		log.Print("unsuccessful setup for user " + test + " test")
		log.Println(err)
		log.Fatal(err)
	}

	u, _ := repo.User.GetByEmail(context.Background(), user.Email)
	return u
}

func UserCleanup(repo *repositories.Repository, user *models.User) {
	if user == nil {
		u, _ := repo.User.GetByEmail(context.Background(), "test-user-create@example.com")
		if u != nil {
			_ = repo.User.Delete(context.Background(), u.UUID)
		}
		u, _ = repo.User.GetByEmail(context.Background(), "test-user-default@example.com")
		if u != nil {
			_ = repo.User.Delete(context.Background(), u.UUID)
		}
		return
	}
	_ = repo.User.Delete(context.Background(), user.UUID)

}
