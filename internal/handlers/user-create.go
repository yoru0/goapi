package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yoru0/goapi.git/internal/models"
)

type UserCreateRequestParam struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateResponseParam struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (h *UserHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	var req UserCreateRequestParam
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
	if len(req.Email) < 5 || !containsAt(req.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	for _, user := range h.users {
		if user.Email == req.Email {
			http.Error(w, "Email already exists", http.StatusConflict)
			return
		}
	}

	req.Name = trimWhitespace(req.Name)
	req.Email = trimWhitespace(req.Email)

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

func trimWhitespace(s string) string {
	start := 0
	end := len(s) - 1

	for start < len(s) && s[start] == ' ' {
		start++
	}

	for end >= 0 && s[end] == ' ' {
		end--
	}

	if start > end {
		return ""
	}

	return s[start : end+1]
}
