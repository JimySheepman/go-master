package domain

import (
	"errors"
	"time"
)

// Order entity
type Order struct {
	ID         string
	CustomerID string
	Amount     float64
	CreatedAt  time.Time
}

// Hatalar
var (
	ErrOrderNotFound = errors.New("order not found")
)
