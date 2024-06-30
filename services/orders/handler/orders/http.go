package handler

import (
	"github.com/emaanmohamed/order-management-system/services/common/genproto/orders"
	"github.com/emaanmohamed/order-management-system/services/orders/types"
	"net/http"
)

type OrderHttpHandler struct {
	orderService types.OrderService
}

func NewOrderHttpHandler(orderService types.OrderService) *OrderHttpHandler {
	return &OrderHttpHandler{
		orderService: orderService,
	}
}

func (h *OrderHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHttpHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	// Create an order object
	// Call the orderService to create the order
	// Write the response
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
