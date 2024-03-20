package sqlstore

import "github.com/XATAB1CH/rest-api/internal/app/model"

// UserRepository
type UserRepository struct {
	store *Store
}

// Создание нового пользователя
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	// Проверка на валидность
	if err := u.Validate(); err != nil {
		return nil, err
	}

	// Хэширование пароля
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	// Запись пользователя в БД
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.Id,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
