package main

func main() {
	httpServer := NewHttpServer(":9003")
	go httpServer.Start()

	grpcServer := NewgRPCServer(":9001")
	grpcServer.Start()
}
