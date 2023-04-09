package main

import (
	"flag"

	"github.com/szymon676/ogauth-grpc/api"
)

func main() {
	port := flag.String("port", ":4000", "port")
	api.NewServer(*port)
}
