package services

import (
	"Hexagonal-Architecture/internal/core/domain"
	"Hexagonal-Architecture/internal/core/ports"
)

type PaymentService struct {
	repo ports.PaymentRepository
}

func NewPaymentService(repo ports.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (p *PaymentService) CreateCheckoutSession(userID string, payment domain.Payment) error {
	return p.repo.CreateCheckoutSession(userID, payment)
}
