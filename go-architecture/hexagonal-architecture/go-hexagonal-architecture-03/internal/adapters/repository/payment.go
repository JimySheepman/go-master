package repository

import "Hexagonal-Architecture/internal/core/domain"

func (c *DB) CreateCheckoutSession(userID string, payment domain.Payment) error {
	return nil
}
