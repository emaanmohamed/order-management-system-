package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return conc

}

func main() {
	httpServer := NewHTTPServer(":8080")
	httpServer.Run()
}
