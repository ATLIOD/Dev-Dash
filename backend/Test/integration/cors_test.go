//go:build integration

package integration

import (
	"DevDash/internal/api"
	"DevDash/internal/api/handlers"
	"DevDash/internal/api/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/jwtauth/v5"
)

func TestCORS(t *testing.T) {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	allowedOrigins := []string{"http://localhost:3000"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders := []string{"Accept", "Authorization", "Content-Type"}
	exposedHeaders := []string{"Link"}

	corsConfig := middleware.GetCorsConfig(allowedOrigins, allowedMethods, allowedHeaders, exposedHeaders)
	h := &handlers.Handler{
		User:    &handlers.UserHandler{},
		Project: &handlers.ProjectHandler{},
	}
	router := api.NewRouter(h, corsConfig, tokenAuth)

	t.Run("Allowed Origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Header().Get("Access-Control-Allow-Origin") != "http://localhost:3000" {
			t.Errorf("expected Access-Control-Allow-Origin: http://localhost:3000, got %s", rr.Header().Get("Access-Control-Allow-Origin"))
		}
	})

	t.Run("Disallowed Origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/health", nil)
		req.Header.Set("Origin", "http://malicious.com")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Header().Get("Access-Control-Allow-Origin") != "" {
			t.Errorf("expected empty Access-Control-Allow-Origin for disallowed origin, got %s", rr.Header().Get("Access-Control-Allow-Origin"))
		}
	})

	t.Run("Preflight Request", func(t *testing.T) {
		req := httptest.NewRequest("OPTIONS", "/health", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		req.Header.Set("Access-Control-Request-Method", "POST")
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
			t.Errorf("expected status 200 OK or 204 NoContent for preflight, got %d", rr.Code)
		}

		if rr.Header().Get("Access-Control-Allow-Origin") != "http://localhost:3000" {
			t.Errorf("expected Access-Control-Allow-Origin: http://localhost:3000, got %s", rr.Header().Get("Access-Control-Allow-Origin"))
		}

		if rr.Header().Get("Access-Control-Allow-Methods") == "" {
			t.Error("expected Access-Control-Allow-Methods header, got empty")
		}
	})
}
