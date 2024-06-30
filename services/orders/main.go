package main

func main() {
	grpcServer := NewGRPCServer("localhost:8050")
	grpcServer.Run()

}
