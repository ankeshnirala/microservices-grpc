package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) (*grpc.ClientConn, error) {
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func main() {
	httpServer := NewHttpServer(":9002")
	log.Fatal(httpServer.Start())
}
