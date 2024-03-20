package teststore

import (
	"github.com/XATAB1CH/rest-api/internal/app/model"
	"github.com/XATAB1CH/rest-api/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

// New
func New() *Store {
	return &Store{}
}

// User
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.userRepository
}
