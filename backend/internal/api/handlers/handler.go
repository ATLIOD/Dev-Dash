package handlers

import (
	"DevDash/internal/services"
	"DevDash/pkg/utils"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	User    *UserHandler
	Project *ProjectHandler
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

func New(svc *services.Service, tokenAuth *jwtauth.JWTAuth) *Handler {
	return &Handler{
		User:    &UserHandler{Service: svc.User, TokenAuth: tokenAuth},
		Project: &ProjectHandler{Service: svc.Project},
	}
}
