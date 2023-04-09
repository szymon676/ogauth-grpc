package validators

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/szymon676/ogauth-grpc/proto"
)

func ValidateRegisterReq(req *proto.RegisterRequest) error {
	if len(req.Username) < 2 {
		return errors.New("please provide a username longer than 2 characters")
	}
	if len(req.Password) < 4 {
		return errors.New("please provide a password longer than 4 characters")
	}
	if len(req.Email) < 5 {
		return errors.New("please provide a valid email address")
	}
	return nil
}

func CorrectReq(req *proto.RegisterRequest) (*proto.RegisterRequest, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	correctedEmail := correctEmail(req.Email)
	return &proto.RegisterRequest{Username: req.Username, Password: hashedPassword, Email: correctedEmail}, nil
}

func correctEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func hashPassword(reqPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
