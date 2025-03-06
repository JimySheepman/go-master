//go:build wireinject
// +build wireinject

package http

import (
	"auth/domain/service"
	"auth/infra/rdb/mysql"
	"auth/presentation/http"
	"auth/usecase"

	"github.com/google/wire"
)

func InitServer() (*http.Server, error) {
	wire.Build(
		http.Set,
		service.Set,
		usecase.Set,
		mysql.Set,
	)
	return nil, nil
}
