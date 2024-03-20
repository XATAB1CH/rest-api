package teststore_test

import (
	"testing"

	"github.com/XATAB1CH/rest-api/internal/app/model"
	"github.com/XATAB1CH/rest-api/internal/store"
	"github.com/XATAB1CH/rest-api/internal/store/teststore"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
