package grpc_server

import (
	"context"
	"log"

	"github.com/bygui86/grpc-samples/go/domain"
)

// server is used to implement helloworld.GreeterServer
type Server struct{}

// SayHello implements service HelloService helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *domain.HelloRequest) (*domain.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &domain.HelloResponse{Greeting: "Hello " + in.Name}, nil
}
