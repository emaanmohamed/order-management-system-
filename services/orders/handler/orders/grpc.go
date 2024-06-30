package handler

import (
	"context"
	"github.com/emaanmohamed/order-management-system/services/common/genproto/orders"
	"github.com/emaanmohamed/order-management-system/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
