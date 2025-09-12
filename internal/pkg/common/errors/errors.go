package errors

import "errors"

// User-related errors
var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidUUID        = errors.New("invalid UUID format")
)
