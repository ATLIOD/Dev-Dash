//go:build integration

package integration

import (
	"DevDash/Test"
	"DevDash/Test/integration/utils"
	"DevDash/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectAPI(t *testing.T) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	var userID string
	repo, cleanup := utils.Setup()
	defer cleanup()

	t.Run("Create Project", func(t *testing.T) {
		utils.UserCleanup(repo, "")
		c := Test.NewChecker(t)
		payload := models.CreateProjectRequest{}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		resp, err := http.Post(baseURL+"/user/", "application/json", bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			respBody, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 201, got %d. Body: %s", resp.StatusCode, string(respBody))
		}
		c.Check(assert.Equal(t, http.StatusCreated, resp.StatusCode, "unexpected response code"))

		var user models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		c.Check(assert.NoError(t, err))

		userID = user.ID
		c.Check(assert.NotEmpty(t, userID, "User ID should not be empty in response"))
		c.Check(assert.Equal(t, payload["name"], user.Name, "name does not match expected value"))
		c.Check(assert.Equal(t, payload["email"], user.Email, "name does not match expected value"))

		if c.Failed() {
			log.Print("User creation test has failed")
		}
		utils.UserCleanup(repo, "")
	})

	t.Run("Get Project", func(t *testing.T) {
		uuid := utils.UserSetup(repo, "Get")
		c := Test.NewChecker(t)
		resp, err := http.Get(baseURL + "/user/" + uuid)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var user models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, uuid, user.ID))

		utils.UserCleanup(repo, user.ID)
	})

	t.Run("Get All Projects For User", func(t *testing.T) {
		uuid := utils.UserSetup(repo, "Get")
		c := Test.NewChecker(t)
		resp, err := http.Get(baseURL + "/user/" + uuid)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var user models.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, uuid, user.ID))

		utils.UserCleanup(repo, user.ID)
	})

	t.Run("Update Project", func(t *testing.T) {
		uuid := utils.UserSetup(repo, "Update")
		c := Test.NewChecker(t)

		payload := map[string]string{
			"name":  "Updated Integration User",
			"email": "test-api-updated@example.com",
		}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		req, err := http.NewRequest(http.MethodPut, baseURL+"/user/"+uuid, bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var user models.ProjectResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, payload["name"], user.Name))
		c.Check(assert.Equal(t, payload["email"], user.Email))

		utils.UserCleanup(repo, uuid)
	})

	t.Run("Delete Project", func(t *testing.T) {
		uuid := utils.UserSetup(repo, "Delete")
		c := Test.NewChecker(t)
		req, err := http.NewRequest(http.MethodDelete, baseURL+"/user/"+uuid, nil)
		c.Check(assert.NoError(t, err))

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		// Verify deletion
		getResp, err := http.Get(baseURL + "/user/" + uuid)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, http.StatusInternalServerError, getResp.StatusCode))

		if c.Failed() {
			utils.UserCleanup(repo, uuid)
		}
	})
}
