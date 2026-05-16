package utils

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
	"log"
)

func ProjectSetup(repo *repositories.Repository, test string) (*models.User, *models.Project) {
	UserCleanup(repo, nil)
	u, _ := repo.User.GetByUUID(context.Background(), UserSetup(repo, "").UUID)
	project := &models.Project{
		UUID:          "2aea5bf5-a98c-4706-af44-9a24852ef1be",
		Name:          "test-project-" + test,
		Description:   "project description",
		Status:        "",
		Stack:         "",
		RepositoryURL: "",
		DeploymentURL: "",
		UserID:        u.ID,
	}

	err := repo.Project.Create(context.Background(), project)
	if err != nil {
		log.Print("unsuccessful setup for project " + test + " test")
		log.Println(err)
		log.Fatal(err)
	}

	return u, project
}

func ProjectCleanup(repo *repositories.Repository, user *models.User) {
	projects, _ := repo.Project.GetAllByUserID(context.Background(), user.ID)
	if projects != nil {
		for _, project := range projects {
			_ = repo.Project.Delete(context.Background(), project.UUID)
		}
	}
	UserCleanup(repo, user)
}
