package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-grpc/domain"
	"github.com/bygui86/go-grpc/greeting-service/grpc_client"
	"github.com/bygui86/go-grpc/kubernetes"
	"github.com/bygui86/go-grpc/logger"
	"github.com/bygui86/go-grpc/utils"

	"google.golang.org/grpc"
)

const (
	serverAddressEnvVar  = "GOGRPC_SERVER_ADDRESS"
	greeetingNameEnvVar  = "GOGRPC_GREETING_NAME"
	kubeProbesNameEnvVar = "GOGRPC_KUBE_PROBES_START"

	serverAddressEnvVarDefault  = "0.0.0.0:50051"
	greeetingNameEnvVarDefault  = "ANONYMOUS"
	kubeProbesNameEnvVarDefault = false
)

func main() {
	serverAddress := utils.GetString(serverAddressEnvVar, serverAddressEnvVarDefault)
	name := utils.GetString(greeetingNameEnvVar, greeetingNameEnvVarDefault)
	kubeProbes := utils.GetBool(kubeProbesNameEnvVar, kubeProbesNameEnvVarDefault)

	grpcConn := createGrpcConnection(serverAddress)
	defer grpcConn.Close()
	logger.SugaredLogger.Infof("gRPC Connection ready to %s", serverAddress)

	go startGreetings(grpcConn, name)

	if kubeProbes {
		kubeServer := startKubernetes(grpcConn)
		defer kubeServer.Shutdown()
	}

	logger.SugaredLogger.Info("Greeting service started!")
	startSysCallChannel()
}

// createGrpcConnection -
func createGrpcConnection(host string) *grpc.ClientConn {
	connection, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		logger.SugaredLogger.Errorf("Connection to gRPC server failed: %v", err.Error())
		os.Exit(3)
	}

	logger.SugaredLogger.Info("State: ", connection.GetState())
	logger.SugaredLogger.Info("Target: ", connection.Target())

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

// startKubernetes -
func startKubernetes(grpcConn *grpc.ClientConn) *kubernetes.KubeProbesServer {
	kubeProbes := kubernetes.KubeProbes{
		GrpcInterface: &grpc_client.GrpcGreetingService{
			GrpcConnection: grpcConn,
		},
	}
	server, err := kubernetes.NewKubeProbesServer(kubeProbes)
	if err != nil {
		logger.SugaredLogger.Errorf("Kubernetes probes server creation failed: %s", err.Error())
		os.Exit(2)
	}
	logger.SugaredLogger.Debug("Kubernetes probes server successfully created")

	server.Start()
	logger.SugaredLogger.Debug("Kubernetes probes successfully started")

	return server
}

// startSysCallChannel -
func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
	logger.SugaredLogger.Info("Termination signal received!")
}
