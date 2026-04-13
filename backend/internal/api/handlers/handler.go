package handlers

import "DevDash/internal/services"

type Handler struct {
	User *UserHandler
}

func New(svc *services.Service) *Handler {
	return &Handler{
		User: &UserHandler{Service: svc.User},
	}
}
