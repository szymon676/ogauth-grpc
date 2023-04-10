package api

import (
	"context"

	"github.com/szymon676/ogauth-grpc/proto"
	"github.com/szymon676/ogauth-grpc/service"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	service *service.UserService
}

func NewAuthServer(service *service.UserService) *AuthServer {
	return &AuthServer{
		service: service,
	}
}

func (s AuthServer) HandleRegister(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	err := s.service.RegisterUser(req)
	if err != nil {
		return &proto.RegisterResponse{
			Message: "user registration failed ;c",
		}, err
	}
	return &proto.RegisterResponse{
		Message: "user registration successfully",
	}, nil
}

func (s AuthServer) HandleLogin(context.Context, *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, nil
}
