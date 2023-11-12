package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/httpserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.GatewayConfig.Level)

	orderConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.OrderConfig.GRPC.Host, cfg.OrderConfig.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create orderConn: %v", err)
	}
	defer orderConn.Close()

	orderClient := pbOrder.NewOrderClient(orderConn)
	gatewayHandler := handler.NewGatewayRoutes(orderClient, logger)

	// Start GRPC Server
	grpcServer, err := grpcserver.New(cfg.GatewayConfig.GRPC.Port)
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create new gateway grpc server: %w", err))
	}
	pbGateway.RegisterGatewayServer(grpcServer.Server, gatewayHandler)
	reflection.Register(grpcServer.Server)

	if err = grpcServer.Start(); err != nil {
		log.Fatalf("failed to start gateway grpcServer: %s", err)
	}
	defer grpcServer.Shutdown()
	logger.Info("GRPC server started")

	// Start HTTP Server
	mux := runtime.NewServeMux()
	err = pbGateway.RegisterGatewayHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("localhost:%v", cfg.GatewayConfig.GRPC.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to register api gateway handler from endpoint: %s", err)
	}

	httpServer := httpserver.New(mux)
	httpServer.Start()
	defer httpServer.Shutdown()
	logger.Info("HTTP server started")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-grpcServer.Notify():
		logger.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	case err = <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
}
