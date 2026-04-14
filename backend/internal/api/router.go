package api

import (
	"DevDash/internal/api/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler) chi.Router {
	r := chi.NewRouter()

	// can uncomment these once middleware is real
	// r.Use(middleware.Cors)
	// r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		h.User.RegisterRoutes(r)
	})

	return r
}
