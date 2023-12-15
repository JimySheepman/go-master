//go:build wireinject
// +build wireinject

package grpc

import (
	"auth/domain/service"
	"auth/infra/rdb/mysql"
	"auth/presentation/grpc"
	"auth/usecase"

	"github.com/google/wire"
)

func InitServer() (*grpc.Server, error) {
	wire.Build(
		grpc.Set,
		service.Set,
		usecase.Set,
		mysql.Set,
	)
	return nil, nil
}
