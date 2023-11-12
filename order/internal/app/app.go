package app

import (
	"fmt"
	"log"
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
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.OrderConfig.Level)

	pg, err := postgres.New(cfg.OrderConfig.PG.URL)
	if err != nil {
		log.Fatalf("failed to run postgres.New: %s", err)
	}
	defer pg.Close()

	orderRepo := repository.NewOrderRepo(pg, logger)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	orderHandler := handler.NewOrderRoutes(*orderUseCase, logger)

	grpcServer, err := grpcserver.New(cfg.OrderConfig.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to create new grcp server: %s", err)
	}
	pbOrder.RegisterOrderServer(grpcServer.Server, orderHandler)
	reflection.Register(grpcServer.Server)

	if err = grpcServer.Start(); err != nil {
		log.Fatalf("failed to start grpcServer: %s", err)
	}
	defer grpcServer.Shutdown()
	logger.Info("GRPC server started")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-grpcServer.Notify():
		logger.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}
}
