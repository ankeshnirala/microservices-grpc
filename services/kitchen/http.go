package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/ankeshnirala/kitchen/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Start() error {
	router := http.NewServeMux()

	conn, err := NewGRPCClient(":9001")
	if err != nil {
		slog.Error("gRPC client failed to connect: %v", err)
	}
	defer conn.Close()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		client := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		_, err = client.CreateOrder(ctx, &orders.CreateOrderRequest{CustomerID: 24, ProductID: 2, Quantity: 5})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		o, err := client.GetOrders(ctx, &orders.GetOrdersRequest{CustomerID: 43})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, o.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	slog.Info(fmt.Sprintf("Starting HTTP server on %s", s.addr))

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
