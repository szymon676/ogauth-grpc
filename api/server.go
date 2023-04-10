package api

import (
	"fmt"
	"log"
	"net"

	"github.com/szymon676/ogauth-grpc/proto"
	"github.com/szymon676/ogauth-grpc/service"
	"github.com/szymon676/ogauth-grpc/store"
	"google.golang.org/grpc"
)

func NewServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("could not listen on port", port)
	}

	store := store.NewScyllaStore()
	service := service.NewUserService(store)
	as := NewAuthServer(service)
	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, as)

	fmt.Println("server started on port", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to reigster server")
	}
}
