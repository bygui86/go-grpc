package grpc_server

import (
	"context"
	"fmt"

	"github.com/bygui86/go-grpc/domain"
	"github.com/bygui86/go-grpc/logger"
)

// Server - Used to implement helloworld.GreeterServer
type Server struct{}

// SayHello - Implement service HelloService helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *domain.HelloRequest) (*domain.HelloResponse, error) {
	logger.SugaredLogger.Infof("Name to greet: %s", in.Name)
	return &domain.HelloResponse{
		Greeting: buildGreet(in.Name),
	}, nil
}

// buildGreet -
func buildGreet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
