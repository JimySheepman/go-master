package service

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/domain"
	"github.com/JimySheepman/go-master/go-algorithm/architecture/onion/internal/repository"
)

// UserService, iş mantığını yöneten servis
type UserService struct {
	repo *repository.InMemoryUserRepository
}

// NewUserService, yeni bir UserService oluşturur
func NewUserService(repo *repository.InMemoryUserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser, yeni kullanıcı oluşturur
func (s *UserService) CreateUser(id, name string) (domain.User, error) {
	user := domain.User{ID: id, Name: name}
	err := s.repo.Save(user)
	return user, err
}

// GetUserByID, kullanıcıyı ID'ye göre getirir
func (s *UserService) GetUserByID(id string) (domain.User, error) {
	return s.repo.FindByID(id)
}
