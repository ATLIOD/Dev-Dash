package handlers

import (
	"DevDash/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service services.UserService
}

// TODO: need to figure out how to actually nest routers a stuff i think

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", h.Get)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
