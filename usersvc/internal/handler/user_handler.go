package handler

import (
	"encoding/json"
	"net/http"

	"github.com/asb19/usersvc/internal/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	service service.UserService
}

func NewHandler(service service.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := uuid.Parse(idStr)
	t, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(t)
}

// GetUsers godoc
// @Summary Get all users
// @Description Returns a list of all users (public info only).
// @Tags users
// @Produce json
// @Success 200 {array} model.UserPublicInfo
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(users)
}
