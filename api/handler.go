package api

import (
	"context"

	"github.com/szymon676/ogauth-grpc/proto"
	"github.com/szymon676/ogauth-grpc/store"
	"github.com/szymon676/ogauth-grpc/validators"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	store store.Store
}

func NewAuthServer(store store.Store) *AuthServer {
	return &AuthServer{
		store: store,
	}
}

func (s AuthServer) HandleRegister(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	err := validators.ValidateRegisterReq(req)
	if err != nil {
		return &proto.RegisterResponse{
			Message: err.Error(),
		}, err
	}
	correctedReq, err := validators.CorrectReq(req)
	if err != nil {
		return &proto.RegisterResponse{
			Message: err.Error(),
		}, err
	}
	err = s.store.SaveUser(correctedReq)
	if err != nil {
		return &proto.RegisterResponse{
			Message: err.Error(),
		}, err
	}
	return &proto.RegisterResponse{
		Message: "user registration successfully",
	}, nil
}

func (s AuthServer) HandleLogin(context.Context, *proto.LoginRequest) (*proto.LoginResponse, error) {
	return nil, nil
}
