package main

import (
	"fmt"
	"log/slog"
	"net"

	handler "github.com/ankeshnirala/kitchen/services/orders/handler/orders"
	"github.com/ankeshnirala/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPC struct {
	addr string
}

func NewgRPCServer(addr string) *gRPC {
	return &gRPC{addr: addr}
}

func (s *gRPC) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		slog.Error("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewOrdersGrpcService(grpcServer, orderService)

	slog.Info(fmt.Sprintf("Starting gRPC server on %s", s.addr))

	return grpcServer.Serve(lis)
}
