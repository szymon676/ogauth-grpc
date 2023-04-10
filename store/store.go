package store

import "github.com/szymon676/ogauth-grpc/types"

type Store interface {
	SaveUser(user types.User) error
}
