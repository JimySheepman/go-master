package persistence

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/ddd/internal/domain"
	"sync"
)

// MemoryOrderRepository, bellek içi sipariş deposu
type MemoryOrderRepository struct {
	orders map[string]domain.Order
	mu     sync.RWMutex
}

// NewMemoryOrderRepository, yeni bir MemoryOrderRepository oluşturur
func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{orders: make(map[string]domain.Order)}
}

// Save, siparişi kaydeder
func (r *MemoryOrderRepository) Save(order domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

// FindByID, siparişi ID'ye göre bulur
func (r *MemoryOrderRepository) FindByID(id string) (domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, exists := r.orders[id]
	if !exists {
		return domain.Order{}, domain.ErrOrderNotFound
	}
	return order, nil
}
