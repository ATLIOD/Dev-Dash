//go:build integration

package integration

import (
	"DevDash/Test"
	"DevDash/Test/integration/utils"
	"DevDash/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectAPI(t *testing.T) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	repo, cleanup := utils.Setup()
	defer cleanup()

	t.Run("Create Project", func(t *testing.T) {
		user := utils.UserSetup(repo, "CreateProject")
		c := Test.NewChecker(t)
		payload := models.CreateProjectRequest{
			Name:          "Test Creation Project",
			Description:   "Test Creation Project Description",
			Status:        "In progress",
			Stack:         "React, Go, Docker",
			RepositoryURL: "example.com",
			DeploymentURL: "example.com",
			UserID:        user.ID,
		}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		resp, err := http.Post(baseURL+"/project/", "application/json", bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusCreated, resp.StatusCode, "unexpected response code"))

		var project models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&project)
		c.Check(assert.NoError(t, err))

		c.Check(assert.NotEmpty(t, project.ID, "Project ID should not be empty in response"))
		c.Check(assert.Equal(t, payload.Name, project.Name, "name does not match expected value"))
		c.Check(assert.Equal(t, payload.UserID, project.UserID, "user id does not match expected value"))

		if c.Failed() {
			log.Print("Project creation test has failed")
		}
		utils.ProjectCleanup(repo, user)
	})

	t.Run("Get Project", func(t *testing.T) {
		user, project := utils.ProjectSetup(repo, "Get")
		c := Test.NewChecker(t)
		resp, err := http.Get(baseURL + "/project/" + project.UUID)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var retrievedProject models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&retrievedProject)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, retrievedProject.ID, project.UUID))
		c.Check(assert.Equal(t, retrievedProject.UserID, user.ID))

		utils.ProjectCleanup(repo, user)
	})

	t.Run("Get All Projects For User", func(t *testing.T) {
		user := utils.UserSetup(repo, "Get")
		for i := range 5 {
			_ = repo.Project.Create(context.Background(), &models.Project{Name: "Temp Project" + strconv.Itoa(i), UserID: user.ID})
		}

		c := Test.NewChecker(t)
		resp, err := http.Get(baseURL + "/projects/" + user.UUID)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var projects []models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&projects)
		c.Check(assert.NoError(t, err))
		c.Check(assert.True(t, len(projects) == 5))

		utils.ProjectCleanup(repo, user)
	})

	t.Run("Update Project", func(t *testing.T) {
		user, project := utils.ProjectSetup(repo, "Update")
		c := Test.NewChecker(t)

		payload := models.UpdateProjectRequest{
			Name:          "Updated Integration Project",
			Description:   "Test Creation Project Description",
			Status:        "In progress",
			Stack:         "React, Go, Docker",
			RepositoryURL: "example.com",
			DeploymentURL: "example.com",
		}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		req, err := http.NewRequest(http.MethodPut, baseURL+"/project/"+project.UUID, bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var updatedProject models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&updatedProject)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, payload.Name, updatedProject.Name))
		c.Check(assert.Equal(t, payload.Description, updatedProject.Description))

		utils.ProjectCleanup(repo, user)
	})

	t.Run("Delete Project", func(t *testing.T) {
		user, project := utils.ProjectSetup(repo, "Delete")
		c := Test.NewChecker(t)
		req, err := http.NewRequest(http.MethodDelete, baseURL+"/project/"+project.UUID, nil)
		c.Check(assert.NoError(t, err))

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		// Verify deletion
		getResp, err := http.Get(baseURL + "/project/" + project.UUID)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, http.StatusInternalServerError, getResp.StatusCode))

		if c.Failed() {
			utils.ProjectCleanup(repo, user)
		}
	})
}
