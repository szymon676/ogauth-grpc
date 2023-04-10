package service

import (
	"github.com/szymon676/ogauth-grpc/proto"
	"github.com/szymon676/ogauth-grpc/store"
	"github.com/szymon676/ogauth-grpc/types"
	"github.com/szymon676/ogauth-grpc/validators"
)

type IUserService interface {
	RegisterUser(req *proto.RegisterRequest) error
}

type UserService struct {
	store store.Store
}

func NewUserService(store store.Store) *UserService {
	return &UserService{store: store}
}

func (s *UserService) RegisterUser(req *proto.RegisterRequest) error {
	err := validators.ValidateRegisterReq(req)
	if err != nil {
		return err
	}
	req, err = validators.CorrectReq(req)
	if err != nil {
		return err
	}

	user := types.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = s.store.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}
