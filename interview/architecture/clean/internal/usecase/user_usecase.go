package usecase

import "github.com/JimySheepman/go-master/go-algorithm/architecture/clean/internal/domain"

// UserUseCase, iş mantığını yöneten usecase
type UserUseCase struct {
	repo domain.UserRepository
}

// NewUserUseCase, yeni bir UserUseCase oluşturur
func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// CreateUser, yeni kullanıcı oluşturur
func (u *UserUseCase) CreateUser(id, name string) (domain.User, error) {
	user := domain.User{ID: id, Name: name}
	err := u.repo.Save(user)
	return user, err
}

// GetUserByID, kullanıcıyı ID'ye göre getirir
func (u *UserUseCase) GetUserByID(id string) (domain.User, error) {
	return u.repo.FindByID(id)
}
