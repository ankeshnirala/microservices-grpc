package handler

import (
	"net/http"

	"github.com/ankeshnirala/kitchen/services/common/genproto/orders"
	"github.com/ankeshnirala/kitchen/services/common/utils"
	"github.com/ankeshnirala/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	orderService types.OrderService
}

func NewOrdersHttpHandler(orderService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{orderService: orderService}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	utils.WriteJSON(w, http.StatusOK, res)
}
