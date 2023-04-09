package store

import (
	"github.com/szymon676/ogauth-grpc/proto"
)

type Store interface {
	SaveUser(user *proto.RegisterRequest) error
}
