package application

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/domain"
	"time"
)

// OrderService, iş mantığını yöneten servis
type OrderService struct {
	orderRepo domain.OrderRepository
}

// NewOrderService, yeni bir OrderService oluşturur
func NewOrderService(orderRepo domain.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

// CreateOrder, yeni sipariş oluşturur
func (s *OrderService) CreateOrder(customerID string, amount float64) (domain.Order, error) {
	order := domain.Order{
		ID:         "random-id", // ID oluşturma işlemi burada yapılmalıdır.
		CustomerID: customerID,
		Amount:     amount,
		CreatedAt:  time.Now(),
	}
	err := s.orderRepo.Save(order)
	return order, err
}

// GetOrderByID, siparişi ID'ye göre getirir
func (s *OrderService) GetOrderByID(id string) (domain.Order, error) {
	return s.orderRepo.FindByID(id)
}
