package application

import (
	"context"
	"go-hexagonal/internal/core/application/dto"
	"go-hexagonal/internal/core/domain"
	outdb "go-hexagonal/internal/core/ports/driven"
	"go-hexagonal/pkg/encryption"
	"go-hexagonal/pkg/uidgen"
	"time"
)

type UserService struct {
	userDB outdb.UserDB
}

func NewUserService(userDB outdb.UserDB) UserService {
	return UserService{
		userDB: userDB,
	}

}

func (usrv UserService) Register(ctx context.Context, user dto.User) (*dto.User, error) {
	id := uidgen.New().New()

	newuser, err := domain.NewUser(id, user.Name, user.Lastname, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	pass, err := encryption.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	passencrypted, _ := domain.NewUserPassword(pass)

	newuser.Password = passencrypted
	now := time.Now()
	newuser.CreatedAt = now
	newuser.UpdatedAt = now

	if err := usrv.userDB.Create(ctx, newuser); err != nil {
		return nil, err
	}

	user.ID = id

	return &user, nil
}
