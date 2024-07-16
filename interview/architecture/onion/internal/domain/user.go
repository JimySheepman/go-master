package domain

import "errors"

// User domain nesnesi
type User struct {
	ID   string
	Name string
}

// Hatalar
var (
	ErrUserNotFound = errors.New("user not found")
)
