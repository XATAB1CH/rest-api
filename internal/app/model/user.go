package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `validate:"email" json:"email"`
	Password          string `validate:"min=6,max=32,compare" json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Custom validation ...
func ComparePasswordValidate(fl validator.FieldLevel) bool {
	u := fl.Parent().Interface().(User)
	if u.EncryptedPassword == "" {
		return true

	}
	return false
}

// Validate ...
func (u *User) Validate() error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("compare", ComparePasswordValidate)

	return validate.Struct(u)
}

// BeforeCreate ...
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

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

// ComparePassword...
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

// Encrypt password ...
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
