package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func TestUserAPI(t *testing.T) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	var userID string

	t.Run("Create User", func(t *testing.T) {
		payload := map[string]string{
			"name":     "Integration Test User",
			"email":    "test-api@example.com",
			"password": "password123",
		}
		body, err := json.Marshal(payload)
		require.NoError(t, err)

		resp, err := http.Post(baseURL+"/user/", "application/json", bytes.NewBuffer(body))
		require.NoError(t, err)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			respBody, _ := io.ReadAll(resp.Body)
			t.Fatalf("Expected status 201, got %d. Body: %s", resp.StatusCode, string(respBody))
		}
		require.Equal(t, http.StatusCreated, resp.StatusCode)

		var user UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		require.NoError(t, err)

		require.Equal(t, "test-api@example.com", user.Email)
		userID = user.ID
		require.NotEmpty(t, userID)
		require.Equal(t, payload["name"], user.Name)
		require.Equal(t, payload["email"], user.Email)
	})

	t.Run("Get User", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/user/" + userID)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var user UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		require.NoError(t, err)
		require.Equal(t, userID, user.ID)
	})

	t.Run("Update User", func(t *testing.T) {
		payload := map[string]string{
			"name":  "Updated Integration User",
			"email": "test-api-updated@example.com",
		}
		body, err := json.Marshal(payload)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, baseURL+"/user/"+userID, bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var user UserResponse
		err = json.NewDecoder(resp.Body).Decode(&user)
		require.NoError(t, err)
		require.Equal(t, payload["name"], user.Name)
		require.Equal(t, payload["email"], user.Email)
	})

	t.Run("Delete User", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, baseURL+"/user/"+userID, nil)
		require.NoError(t, err)

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		// Verify deletion
		getResp, err := http.Get(baseURL + "/user/" + userID)
		require.NoError(t, err)
		require.Equal(t, http.StatusInternalServerError, getResp.StatusCode)
	})
}
