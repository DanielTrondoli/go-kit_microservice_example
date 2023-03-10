package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndPoints(s Service) EndPoints {
	return EndPoints{
		CreateUser: makeCreateUserEndPoint(s),
		GetUser:    makeGetUserEndPoint(s),
	}
}

func makeCreateUserEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(ctx, req.Email, req.Password)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		email, err := s.GetUser(ctx, req.Id)
		return GetUserResponse{Email: email}, err
	}
}
