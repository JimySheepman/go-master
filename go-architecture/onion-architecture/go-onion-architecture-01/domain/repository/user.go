//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_$GOPACKAGE/$GOFILE

package repository

import "go-onion-architecture-sample/domain/model"

type UserRepository interface {
	Create(string) (model.UserId, error)
	Get(model.UserId) (model.User, error)
}
