package api

import (
	"DevDash/internal/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func NewRouter(h *handlers.Handler, corsConfig *cors.Cors, tokenAuth *jwtauth.JWTAuth) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsConfig.Handler)

	r.Get("/health", h.Health)

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/user/", h.User.Create)
		r.Post("/user/login", h.User.Login)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		h.User.RegisterRoutes(r)
		h.Project.RegisterRoutes(r)
	})

	return r
}
