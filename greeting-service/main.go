package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/bygui86/grpc-samples/go/domain"

	"google.golang.org/grpc"
)

const (
	defaultGrpcServerAddress = "localhost:50051"
	defaultName              = "ANONYMOUS"
)

var (
	grpcServerAddress string
	name              string
)

func init() {
	flag.StringVar(&grpcServerAddress, "grpcServerAddress", defaultGrpcServerAddress, "Server host:port")
	flag.StringVar(&name, "name", defaultName, "Name to greet")
	flag.Parse()
}

func main() {
	// Set up a connection to the server.
	connection, connErr := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if connErr != nil {
		log.Fatalf("Connaction to gRPC server failed: %v", connErr)
		panic(connErr)
	}
	defer connection.Close()
	client := domain.NewHelloServiceClient(connection)

	// Define timeout
	timeout := 1 * time.Second

	// Contact the server
	// WARNING: the connection context is one-shot, it must be refreshed before every request
	connectionCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	response, reqErr := client.SayHello(connectionCtx, &domain.HelloRequest{Name: name})
	if reqErr != nil {
		log.Fatalf("could not greet: %v", reqErr)
		panic(reqErr)
	}

	// Print out the response
	log.Printf("Greeting: %s", response.Greeting)
}
