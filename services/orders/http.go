package main

import (
	handler "github.com/emaanmohamed/order-management-system/services/orders/handler/orders"
	"github.com/emaanmohamed/order-management-system/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct {
	addr string
}

func NewHTTPServer() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrderHttpHandler(orderService)

	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
