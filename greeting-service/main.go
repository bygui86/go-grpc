package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-grpc/domain"
	"github.com/bygui86/go-grpc/logger"
	"github.com/bygui86/go-grpc/utils"

	"google.golang.org/grpc"
)

const (
	serverAddressEnvVar = "GOGERPC_GRPC_SERVER_ADDRESS"
	greeetingNameEnvVar = "GOGERPC_GREETING_NAME"

	serverAddressEnvVarDefault = "0.0.0.0:50051"
	greeetingNameEnvVarDefault = "ANONYMOUS"
)

func main() {
	serverAddress := utils.GetString(serverAddressEnvVar, serverAddressEnvVarDefault)
	name := utils.GetString(greeetingNameEnvVar, greeetingNameEnvVarDefault)

	connection := createGrpcConnection(serverAddress)
	defer connection.Close()
	logger.SugaredLogger.Infof("Connection ready to %s", serverAddress)

	go startGreetings(connection, name)

	logger.SugaredLogger.Info("Greeting service started!")
	startSysCallChannel()
}

// createGrpcConnection -
func createGrpcConnection(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		logger.SugaredLogger.Errorf("Connection to gRPC server failed: %v", err.Error())
		os.Exit(2)
	}
	return connection
}

// startGreetings -
func startGreetings(connection *grpc.ClientConn, name string) {
	timeout := 2 * time.Second
	client := domain.NewHelloServiceClient(connection)
	logger.SugaredLogger.Infof("Starting greeting %s...", name)
	for {
		go greet(client, timeout, name)
		time.Sleep(3 * time.Second)
	}
}

// greet -
func greet(client domain.HelloServiceClient, timeout time.Duration, name string) {
	// WARNING: the connection context is one-shot, it must be refreshed before every request
	connectionCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	response, err := client.SayHello(connectionCtx, &domain.HelloRequest{Name: name})
	if err != nil {
		logger.SugaredLogger.Errorf("Could not greet %s: %v", name, err.Error())
		return
	}
	logger.SugaredLogger.Info(response.Greeting)
}

// startSysCallChannel -
func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}
