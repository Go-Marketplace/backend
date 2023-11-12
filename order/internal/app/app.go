package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/order/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/order/internal/infrastructure/repository"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	pbOrder "github.com/Go-Marketplace/backend/proto/order/gen/proto"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.OrderConfig.Level)

	pg, err := postgres.New(cfg.OrderConfig.PG.URL)
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to run postgres.New: %w", err))
	}
	defer pg.Close()

	orderRepo := repository.NewOrderRepo(pg, logger)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	orderHandler := handler.NewOrderRoutes(*orderUseCase, logger)

	grpcServer, err := grpcserver.New(cfg.OrderConfig.GRPC.Port)
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create new grcp server: %w", err))
	}
	pbOrder.RegisterOrderServer(grpcServer.Server, orderHandler)

	if err = grpcServer.Start(); err != nil {
		logger.Fatal(fmt.Errorf("failed to start grpcServer: %w", err))
	}
	defer grpcServer.Shutdown()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-grpcServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}
}
