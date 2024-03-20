package teststore

import (
	"github.com/XATAB1CH/rest-api/internal/app/model"
	"github.com/XATAB1CH/rest-api/internal/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {

	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FIndByEmail
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return u, nil
}
