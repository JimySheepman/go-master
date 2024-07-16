package repository

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/hexagonal/domain"
	"sync"
)

// InMemoryUserRepository, bellek içi kullanıcı deposu
type InMemoryUserRepository struct {
	users map[string]domain.User
	mu    sync.RWMutex
}

// NewInMemoryUserRepository, yeni bir InMemoryUserRepository oluşturur
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]domain.User)}
}

// Save, kullanıcıyı kaydeder
func (r *InMemoryUserRepository) Save(user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

// FindByID, kullanıcıyı ID'ye göre bulur
func (r *InMemoryUserRepository) FindByID(id string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[id]
	if !exists {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}
