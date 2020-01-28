package main

import (
	"log"
	"net"

	"github.com/bygui86/grpc-samples/go/domain"
	"github.com/bygui86/grpc-samples/go/hello-service/grpc_server"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// TCP listener
	listener, tcpErr := net.Listen("tcp", port)
	if tcpErr != nil {
		log.Fatalf("Failed to listen: %v", tcpErr)
		panic(tcpErr)
	}
	log.Println("TCP listener ready on port ", port)

	// gRPC server
	grpcServer := grpc.NewServer()
	helloSvcServer := grpc_server.Server{}
	domain.RegisterHelloServiceServer(grpcServer, &helloSvcServer)
	if grpcErr := grpcServer.Serve(listener); grpcErr != nil {
		log.Fatalf("Failed to serve: %v", grpcErr)
		panic(grpcErr)
	}
}
