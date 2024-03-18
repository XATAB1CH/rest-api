package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

// User
type User struct {
	Id                int
	Email             string `validate:"email"`
	Password          string `validate:"custom"`
	EncryptedPassword string
}

func customFunc(fl validator.FieldLevel) bool {
	u := fl.Field().Interface().(User)

	if u.EncryptedPassword != "" {
		return true
	}

	return false
}

// Validate
func (u *User) Validate() error {
	validate.RegisterValidation("custom", customFunc)
	validate = validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(u)
}

// BeforeCreate
func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// Encrypt password
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
