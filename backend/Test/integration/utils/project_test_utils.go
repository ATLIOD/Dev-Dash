package utils

import (
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"context"
	"log"
)

func ProjectSetup(repo *repositories.Repository, test string) string {
	u, _ := repo.User.GetByUUID(context.Background(), UserSetup(repo, "project"))
	project := &models.Project{
		ID:            1,
		UUID:          "01",
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

	v, _ := repo.Project.GetByEmail(context.Background(), project.Email)
	return v.UUID
}

func ProjectCleanup(repo *repositories.Repository, uuid string) {
	if uuid == "" {
		project, _ := repo.Project.GetByEmail(context.Background(), "test-project-create@example.com")
		if project != nil {
			_ = repo.Project.Delete(context.Background(), project.UUID)
		}
		return
	}
	_ = repo.Project.Delete(context.Background(), uuid)

}
