package mock

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/yoru0/goapi.git/internal/pkg/models"
)

type UserStore struct {
	mu    sync.RWMutex
	users map[string]*models.User
}

func NewUserStore() *UserStore {
	store := &UserStore{
		users: make(map[string]*models.User),
	}

	store.seedDummy()
	return store
}

func (s *UserStore) seedDummy() {
	dummy := []*models.User{
		{
			ID:        uuid.New().String(),
			Name:      "Jes",
			Email:     "jes@example.com",
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			Name:      "jor",
			Email:     "jor.smith@example.com",
			CreatedAt: time.Now().Add(-12 * time.Hour),
			UpdatedAt: time.Now().Add(-12 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			Name:      "shama",
			Email:     "shama@example.com",
			CreatedAt: time.Now().Add(-6 * time.Hour),
			UpdatedAt: time.Now().Add(-6 * time.Hour),
		},
	}

	for _, user := range dummy {
		s.users[user.ID] = user
	}
}

// Store adds a new user to the store and assigns a unique ID.
func (s *UserStore) Store(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = uuid.New().String()
	s.users[user.ID] = user
}

// FindByID retrieves a user by ID.
func (s *UserStore) FindByID(id string) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[id]
	return user, exists
}

// FindByEmail retrieves a user by email.
func (s *UserStore) FindByEmail(email string) (*models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, user := range s.users {
		if user.Email == email {
			return user, true
		}
	}
	return nil, false
}

// FindAll retrieves all users in the store.
func (s *UserStore) FindAll() []*models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	users := make([]*models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// Update modifies an existing user.
func (s *UserStore) Update(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[user.ID] = user
}

// Remove deletes a user by ID.
func (s *UserStore) Remove(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.users[id]
	if exists {
		delete(s.users, id)
	}
	return exists
}
