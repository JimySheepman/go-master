package service

import (
	"mastering-grpc/internal/core/model/request"
	"mastering-grpc/internal/core/model/response"
)

type UserService interface {
	SignUp(request *request.SignUpRequest) *response.Response
}
