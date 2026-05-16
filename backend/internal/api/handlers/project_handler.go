package handlers

import (
	"DevDash/internal/models"
	"DevDash/internal/services"
	"DevDash/pkg/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProjectHandler struct {
	Service services.ProjectService
}

func (h *ProjectHandler) RegisterRoutes(r chi.Router) {
	r.Route("/project", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Route("/{projectID}", func(r chi.Router) {
			r.Get("/", h.Get)
			r.Put("/", h.Update)
			r.Delete("/", h.Delete)
		})
	})
	r.Route("/projects", func(r chi.Router) {
		r.Route("/{userUUID}", func(r chi.Router) {
			r.Get("/", h.GetAll)
		})
	})
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateProjectRequest
	err := utils.DecodeJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	project, err := h.Service.Create(r.Context(), req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusCreated, project)
}
func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	if projectID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid project ID")
		return
	}

	project, err := h.Service.GetByUUID(r.Context(), projectID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, project)
}

func (h *ProjectHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userUUID := chi.URLParam(r, "userUUID")
	if userUUID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	projects, err := h.Service.GetAllByUserUUID(r.Context(), userUUID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, projects)
}

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateProjectRequest
	projectID := chi.URLParam(r, "projectID")
	if projectID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid project ID")
		return
	}
	err := utils.DecodeJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	project, err := h.Service.Update(r.Context(), projectID, req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, project)
}
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	if projectID == "" {
		utils.WriteError(w, http.StatusBadRequest, "invalid project ID")
		return
	}
	err := h.Service.Delete(r.Context(), projectID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
