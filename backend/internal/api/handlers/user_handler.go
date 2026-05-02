package handlers

import (
	"DevDash/internal/models"
	"DevDash/internal/services"
	"DevDash/pkg/utils"
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

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	err := utils.DecodeJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	user, err := h.Service.Create(r.Context(), req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusCreated, user)
}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	user, err := h.Service.GetByID(r.Context(), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateUserRequest
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid user ID")
		return
	}
	err := utils.DecodeJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	user, err := h.Service.Update(r.Context(), userID, req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid user ID")
		return
	}
	err := h.Service.Delete(r.Context(), userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
