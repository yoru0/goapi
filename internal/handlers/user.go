package handlers

import (
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

// Helper function to validate email format (simple check).
func containsAt(email string) bool {
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}
