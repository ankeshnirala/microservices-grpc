package service

import (
	"context"

	"github.com/ankeshnirala/kitchen/services/common/genproto/orders"
)

var ordersStorage = make([]*orders.Order, 0)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersStorage = append(ordersStorage, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersStorage
}
