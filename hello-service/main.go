package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bygui86/go-grpc/domain"
	"github.com/bygui86/go-grpc/hello-service/grpc_server"
	"github.com/bygui86/go-grpc/logger"
	"github.com/bygui86/go-grpc/utils"

	"google.golang.org/grpc"
)

const (
	portEnvVar = "GOGRPC_GRPC_SERVER_ADDRESS"

	portEnvVarDefault = "0.0.0.0:50051"
)

func main() {
	address := utils.GetString(portEnvVar, portEnvVarDefault)

	listener := createTcpListener(address)
	defer listener.Close()
	logger.SugaredLogger.Infof("TCP listener ready on %s", address)

	go startGrpcServer(listener)
	logger.SugaredLogger.Infof("gRPC server ready")

	logger.SugaredLogger.Info("Hello service started!")
	startSysCallChannel()
}

// createTcpListener -
func createTcpListener(port string) net.Listener {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.SugaredLogger.Errorf("Failed to listen: %v", err.Error())
		os.Exit(2)
	}
	return listener
}

// startGrpcServer -
func startGrpcServer(listener net.Listener) {
	grpcServer := grpc.NewServer()
	helloSvcServer := grpc_server.Server{}
	domain.RegisterHelloServiceServer(grpcServer, &helloSvcServer)
	err := grpcServer.Serve(listener)
	if err != nil {
		logger.SugaredLogger.Errorf("Failed to serve: %v", err)
		os.Exit(3)
	}
}

// startSysCallChannel -
func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}
