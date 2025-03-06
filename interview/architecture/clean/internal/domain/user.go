package domain

import "errors"

// User domain nesnesi
type User struct {
	ID   string
	Name string
}

// UserRepository arayüzü
type UserRepository interface {
	Save(user User) error
	FindByID(id string) (User, error)
}

// Hatalar
var (
	ErrUserNotFound = errors.New("user not found")
)
