package ports

import (
	"github.com/JimySheepman/go-master/go-algorithm/architecture/hexagonal/domain"
)

// UserRepository, kullanıcı verilerini yöneten port
type UserRepository interface {
	Save(user domain.User) error
	FindByID(id string) (domain.User, error)
}
