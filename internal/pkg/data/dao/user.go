package dao

import (
	"time"

	"github.com/google/uuid"
	"github.com/yoru0/goapi.git/internal/pkg/common/errors"
	"github.com/yoru0/goapi.git/internal/pkg/data/mock"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

var defaultUserDAO = mock.NewUserStore()

type UserDAOInterface interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(id string, updates *models.User) (*models.User, error)
	Delete(id string) error
}

type UserDAO struct {
	store *mock.UserStore
}

func NewUserDAO() UserDAOInterface {
	return &UserDAO{
		store: defaultUserDAO,
	}
}

// Create adds a new user to the store.
func (dao *UserDAO) Create(user *models.User) (*models.User, error) {
	if _, exists := dao.store.FindByEmail(user.Email); exists {
		return nil, errors.ErrEmailAlreadyExists
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	dao.store.Store(user)
	return user, nil
}

// GetByID retrieves a user by their unique ID.
func (dao *UserDAO) GetByID(id string) (*models.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.ErrInvalidUUID
	}

	user, exists := dao.store.FindByID(id)
	if !exists {
		return nil, errors.ErrUserNotFound
	}
	return user, nil
}

// GetByEmail retrieves a user by their email address.
func (dao *UserDAO) GetByEmail(email string) (*models.User, error) {
	user, exists := dao.store.FindByEmail(email)
	if !exists {
		return nil, errors.ErrUserNotFound
	}
	return user, nil
}

// GetAll retrieves all users from the store.
func (dao *UserDAO) GetAll() ([]*models.User, error) {
	users := dao.store.FindAll()
	return users, nil
}

// Update modifies an existing user's details.
func (dao *UserDAO) Update(id string, updates *models.User) (*models.User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errors.ErrInvalidUUID
	}

	user, exists := dao.store.FindByID(id)
	if !exists {
		return nil, errors.ErrUserNotFound
	}

	if updates.Email != "" && updates.Email != user.Email {
		if _, exists := dao.store.FindByEmail(updates.Email); exists {
			return nil, errors.ErrEmailAlreadyExists
		}
		user.Email = updates.Email
	}

	if updates.Name != "" {
		user.Name = updates.Name
	}

	user.UpdatedAt = time.Now()
	dao.store.Update(user)
	return user, nil
}

// Delete removes a user from the store by their ID.
func (dao *UserDAO) Delete(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.ErrInvalidUUID
	}

	if !dao.store.Remove(id) {
		return errors.ErrUserNotFound
	}
	return nil
}
