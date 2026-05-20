package api

import (
	"DevDash/internal/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(h *handlers.Handler, corsConfig *cors.Cors) chi.Router {
	r := chi.NewRouter()

	// can uncomment these once middleware is real
	r.Use(corsConfig.Handler)
	// r.Use(middleware.Logger)

	r.Get("/health", h.Health)
	r.Route("/", func(r chi.Router) {
		h.User.RegisterRoutes(r)
		h.Project.RegisterRoutes(r)
	})

	return r
}
