//go:build integration

package Unit

import (
	"DevDash/Test"
	"DevDash/Test/Unit/utils"
	"DevDash/internal/models"
	"DevDash/internal/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService(t *testing.T) {
	repo := utils.Setup()
	svc := services.New(repo)

	t.Run("Create Project", func(t *testing.T) {
		c := Test.NewChecker(t)
		req := models.CreateProjectRequest{
			Name:          "New Project",
			Description:   "Description",
			Status:        "active",
			Stack:         "go",
			RepositoryURL: "http://github.com",
			DeploymentURL: "http://deploy.com",
			UserID:        1,
		}

		resp, err := svc.Project.Create(context.Background(), req)
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, req.Name, resp.Name))
	})

	t.Run("Get Project", func(t *testing.T) {
		c := Test.NewChecker(t)
		resp, err := svc.Project.GetByUUID(context.Background(), "01")
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, "Project 1", resp.Name))
	})

	t.Run("Update Project", func(t *testing.T) {
		c := Test.NewChecker(t)
		req := models.UpdateProjectRequest{
			Name:        "Updated Project",
			Description: "Updated Description",
		}
		resp, err := svc.Project.Update(context.Background(), "01", req)
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, req.Name, resp.Name))
	})

	t.Run("Delete Project", func(t *testing.T) {
		c := Test.NewChecker(t)
		err := svc.Project.Delete(context.Background(), "01")
		c.Check(assert.NoError(t, err))

		_, err = svc.Project.GetByUUID(context.Background(), "01")
		c.Check(assert.Error(t, err))
		c.Check(assert.Contains(t, err.Error(), "no project found"))
	})
}
