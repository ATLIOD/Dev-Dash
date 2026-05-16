//go:build integration

package integration

import (
	"DevDash/Test"
	"DevDash/Test/integration/utils"
	"DevDash/internal/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAPI(t *testing.T) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	var userID string
	repo, cleanup := utils.Setup()
	defer cleanup()

	t.Run("Create User", func(t *testing.T) {
		utils.UserCleanup(repo, nil)
		c := Test.NewChecker(t)
		payload := models.CreateUserRequest{
			Name:     "test-user-create",
			Email:    "test-user-create@example.com",
			Password: "password123",
		}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		resp, err := http.Post(baseURL+"/user/", "application/json", bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusCreated, resp.StatusCode, "unexpected response code"))

		var user models.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		c.Check(assert.NoError(t, err))

		userID = user.ID
		c.Check(assert.NotEmpty(t, userID, "User ID should not be empty in response"))
		c.Check(assert.Equal(t, payload.Name, user.Name, "name does not match expected value"))
		c.Check(assert.Equal(t, payload.Email, user.Email, "email does not match expected value"))

		if c.Failed() {
			log.Print("User creation test has failed")
		}
		utils.UserCleanup(repo, nil)
	})

	t.Run("Get User", func(t *testing.T) {
		user := utils.UserSetup(repo, "Get")
		c := Test.NewChecker(t)
		resp, err := http.Get(baseURL + "/user/" + user.UUID)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var retrievedUser models.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&retrievedUser)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, user.UUID, retrievedUser.ID))

		utils.UserCleanup(repo, user)
	})

	t.Run("Update User", func(t *testing.T) {
		user := utils.UserSetup(repo, "Update")
		c := Test.NewChecker(t)

		payload := models.UpdateUserRequest{
			Name:  "Updated Integration User",
			Email: "test-api-updated@example.com",
		}
		body, err := json.Marshal(payload)
		c.Check(assert.NoError(t, err))

		req, err := http.NewRequest(http.MethodPut, baseURL+"/user/"+user.UUID, bytes.NewBuffer(body))
		c.Check(assert.NoError(t, err))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		var updatedUser models.UserResponse
		err = json.NewDecoder(resp.Body).Decode(&updatedUser)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, payload.Name, updatedUser.Name))
		c.Check(assert.Equal(t, payload.Email, updatedUser.Email))

		utils.UserCleanup(repo, user)
	})

	t.Run("Delete User", func(t *testing.T) {
		user := utils.UserSetup(repo, "Delete")
		c := Test.NewChecker(t)
		req, err := http.NewRequest(http.MethodDelete, baseURL+"/user/"+user.UUID, nil)
		c.Check(assert.NoError(t, err))

		client := &http.Client{}
		resp, err := client.Do(req)
		c.Check(assert.NoError(t, err))
		defer resp.Body.Close()

		c.Check(assert.Equal(t, http.StatusOK, resp.StatusCode))

		// Verify deletion
		getResp, err := http.Get(baseURL + "/user/" + user.UUID)
		c.Check(assert.NoError(t, err))
		c.Check(assert.Equal(t, http.StatusInternalServerError, getResp.StatusCode))

		if c.Failed() {
			utils.UserCleanup(repo, user)
		}
	})
}
