package handlers

import (
	"DevDash/internal/models"
	"DevDash/internal/services"
	"DevDash/pkg/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type UserHandler struct {
	Service   services.UserService
	TokenAuth *jwtauth.JWTAuth
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Post("/login", h.Login)
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

	user, err := h.Service.GetByUUID(r.Context(), userID)
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

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	user, err := h.Service.Authenticate(r.Context(), req.Email, req.Password)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateToken(h.TokenAuth, user.UUID, 24*time.Hour)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "could not generate token")
		return
	}

	utils.WriteJSON(w, http.StatusOK, models.LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}
