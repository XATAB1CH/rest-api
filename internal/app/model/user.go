package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

// User
type User struct {
	ID                int    `json:"id"`
	Email             string `validate:"email" json:"email"`
	Password          string `validate:"min=6,max=32,password"`
	EncryptedPassword string
}

// Custom validation
func PasswordValidate(fl validator.FieldLevel) bool {
	u := fl.Parent().Interface().(User) //get user
	if u.EncryptedPassword == "" {
		return false
	}
	return true
}

// Validate
func (u *User) Validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("password", PasswordValidate) ///register custom validation

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
