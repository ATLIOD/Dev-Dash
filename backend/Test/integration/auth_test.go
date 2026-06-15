//go:build integration

package integration

import (
	"DevDash/internal/api"
	"DevDash/internal/api/handlers"
	"DevDash/internal/api/middleware"
	"DevDash/internal/models"
	"DevDash/internal/repositories"
	"DevDash/internal/services"
	"DevDash/pkg/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	// 1. Setup Mock Environment
	jwtSecret := "test-secret"
	tokenAuth := jwtauth.New("HS256", []byte(jwtSecret), nil)

	mockDB := models.NewMockDB()
	// Seed a user with a known password for testing login
	user1 := mockDB.Users["01"]
	hashedPassword, _ := utils.HashPassword("password123")
	user1.PasswordHash = hashedPassword
	mockDB.Users["01"] = user1

	repo := repositories.NewMockRepo(mockDB)
	svc := services.New(repo)
	h := handlers.New(svc, tokenAuth)

	corsConfig := middleware.GetCorsConfig([]string{"*"}, []string{"GET", "POST"}, []string{"Content-Type"}, []string{})
	router := api.NewRouter(h, corsConfig, tokenAuth)

	t.Run("Successful Login", func(t *testing.T) {
		loginReq := models.LoginRequest{
			Email:    "user1@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(loginReq)
		req := httptest.NewRequest("POST", "/user/login", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var resp models.LoginResponse
		err := json.NewDecoder(rr.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Token)
		assert.Equal(t, "user1@example.com", resp.User.Email)
	})

	t.Run("Protected Route - Unauthorized", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user/01", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		// Should fail because no token was provided
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("Protected Route - Authorized", func(t *testing.T) {
		// Manually generate a valid token for "user 01"
		token, _ := utils.GenerateToken(tokenAuth, "01", 1*time.Hour)

		req := httptest.NewRequest("GET", "/user/01", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		// Should succeed because a valid token was provided
		assert.Equal(t, http.StatusOK, rr.Code)

		var user models.UserResponse
		err := json.NewDecoder(rr.Body).Decode(&user)
		assert.NoError(t, err)
		assert.Equal(t, "01", user.UUID)
	})
}
