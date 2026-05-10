package handlers

import (
	"DevDash/internal/services"
	"DevDash/pkg/utils"
	"net/http"
)

type Handler struct {
	User    *UserHandler
	Project *ProjectHandler
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

func New(svc *services.Service) *Handler {
	return &Handler{
		User:    &UserHandler{Service: svc.User},
		Project: &ProjectHandler{Service: svc.Project},
	}
}
