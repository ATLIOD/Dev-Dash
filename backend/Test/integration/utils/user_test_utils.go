package utils

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
	"log"
)

func UserSetup(repo *repositories.Repository, test string) string {
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

	v, _ := repo.User.GetByEmail(context.Background(), user.Email)
	return v.UUID
}

func UserCleanup(repo *repositories.Repository, uuid string) {
	if uuid == "" {
		user, _ := repo.User.GetByEmail(context.Background(), "test-user-create@example.com")
		if user != nil {
			_ = repo.User.Delete(context.Background(), user.UUID)
		}
		return
	}
	_ = repo.User.Delete(context.Background(), uuid)

}
