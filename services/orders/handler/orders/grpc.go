package handler

import (
	"context"

	"github.com/ankeshnirala/kitchen/services/common/genproto/orders"
	"github.com/ankeshnirala/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersGrpcService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	o := h.orderService.GetOrders(ctx)

	resp := &orders.GetOrdersResponse{
		Orders: o,
	}

	return resp, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    43,
		CustomerID: 56,
		ProductID:  2,
		Quantity:   10,
	}

	if err := h.orderService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	resp := &orders.CreateOrderResponse{
		Status: "success",
	}

	return resp, nil
}
