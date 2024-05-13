package main

import (
	"fmt"
	"log/slog"
	"net/http"

	handler "github.com/ankeshnirala/kitchen/services/orders/handler/orders"
	"github.com/ankeshnirala/kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Start() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	handler.NewOrdersHttpHandler(orderService).RegisterRouter(router)

	slog.Info(fmt.Sprintf("Starting HTTP server on %s", s.addr))

	return http.ListenAndServe(s.addr, router)
}
