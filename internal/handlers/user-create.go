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
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *models.User `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

// UserCreate handles POST /api/v1/users/create - Create a new user.
func (h *UserHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req UserCreateRequestParam
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := UserCreateResponseParam{
			Success: false,
			Error:   "Invalid JSON format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if req.Name == "" {
		response := UserCreateResponseParam{
			Success: false,
			Error:   "Name is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if req.Email == "" {
		response := UserCreateResponseParam{
			Success: false,
			Error:   "Email is required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if len(req.Email) < 5 || !containsAt(req.Email) {
		response := UserCreateResponseParam{
			Success: false,
			Error:   "Invalid email format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	req.Name = trimWhitespace(req.Name)
	req.Email = trimWhitespace(req.Email)

	if req.Name == "" {
		response := UserCreateResponseParam{
			Success: false,
			Error:   "Name cannot be empty or just spaces",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, user := range h.users {
		if user.Email == req.Email {
			response := UserCreateResponseParam{
				Success: false,
				Error:   "Email already exists",
			}
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	user := models.User{
		ID:        h.nextID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	h.users = append(h.users, user)
	h.nextID++

	response := UserCreateResponseParam{
		Success: true,
		Message: "User created successfully",
		Data:    &user,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
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
