package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yoru0/goapi.git/internal/models"
)

// UserHandler handles user HTTP requests.
type UserHandler struct {
	users  []models.User
	nextID int64
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler() *UserHandler {
	return &UserHandler{
		users:  make([]models.User, 0),
		nextID: 1,
	}
}

// GetUsers handles GET /users - returns all users.
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}

// GetUser handles GET /users/{id} - returns a user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for _, user := range h.users {
		if user.ID == int64(id) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// CreateUser handles POST /users - creates a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	user := models.User{
		ID:        h.nextID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	h.users = append(h.users, user)
	h.nextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser handles DELETE /users/{id} - deletes a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, user := range h.users {
		if user.ID == int64(id) {
			h.users = append(h.users[:i], h.users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func (h *UserHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
