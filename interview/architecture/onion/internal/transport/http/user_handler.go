package http

import (
	"encoding/json"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/domain"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/service"
	"net/http"
)

// UserHandler, HTTP isteklerini yöneten handler
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler, yeni bir UserHandler oluşturur
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser, HTTP POST isteğini işler
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.CreateUser(user.ID, user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// GetUserByID, HTTP GET isteğini işler
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUserByID(id)
	if err != nil {
		if err == domain.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(user)
}
