package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/Go-Marketplace/backend/user/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/user/internal/infrastructure/repository"
	"github.com/Go-Marketplace/backend/user/internal/usecase"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.UserConfig.Level)

	pg, err := postgres.New(cfg.UserConfig.PG.URL)
	if err != nil {
		log.Fatalf("failed to run postgres.New: %s", err)
	}
	defer pg.Close()

	userRepo := repository.NewUserRepo(pg, logger)
	userUsecase := usecase.NewUserUsecase(userRepo, logger)
	userHandler := handler.NewUserRoutes(userUsecase, logger)

	grpcServer, err := grpcserver.New(cfg.UserConfig.GRPC.Port)
	if err != nil {
		log.Fatalf("failed to create new grcp server: %s", err)
	}
	pbUser.RegisterUserServer(grpcServer.Server, userHandler)
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
