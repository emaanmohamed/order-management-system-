package main

func main() {

	httpServer := NewHTTPServer()
	go httpServer.Run()

	grpcServer := NewGRPCServer("localhost:8050")
	grpcServer.Run()

}
