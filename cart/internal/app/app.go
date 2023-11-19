package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-Marketplace/backend/cart/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/cart/internal/api/grpc/interceptors"
	"github.com/Go-Marketplace/backend/cart/internal/infrastructure/repository"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	"github.com/Go-Marketplace/backend/cart/internal/worker"
	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/Go-Marketplace/backend/pkg/redis"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.CartConfig.Level)

	ctx := context.Background()

	redis, err := redis.New(ctx, cfg.CartConfig.Redis.Url)
	if err != nil {
		log.Fatalf("failed to run redis.New: %s", err)
	}
	defer redis.Close()

	pg, err := postgres.New(cfg.CartConfig.PG.URL)
	if err != nil {
		log.Fatalf("failed to run postgres.New: %s", err)
	}
	defer pg.Close()

	cartRepo := repository.NewCartRepo(pg, logger)
	cartTaskRepo := repository.NewCartTaskRepo(redis, logger)
	cartUsecase := usecase.NewCartUsecase(cartRepo, cartTaskRepo, logger)
	cartHandler := handler.NewCartRoutes(cartUsecase, logger)

	interceptor := interceptors.NewInterceptorManager(logger)
	grpcServer, err := grpcserver.New(
		cfg.CartConfig.GRPC.Port,
		grpc.UnaryInterceptor(
			interceptor.LogRequest,
		),
	)
	if err != nil {
		log.Fatalf("failed to create new grcp server: %s", err)
	}
	pbCart.RegisterCartServer(grpcServer.Server, cartHandler)
	reflection.Register(grpcServer.Server)

	if err = grpcServer.Start(); err != nil {
		log.Fatalf("failed to start grpcServer: %s", err)
	}
	defer grpcServer.Shutdown()
	logger.Info("GRPC server started")

	cartTaskWorker := worker.NewCartTaskWorker(cartTaskRepo, cartUsecase, logger)
	cartTaskWorker.Run(ctx)
	defer func () {
		if err := cartTaskWorker.Stop(); err != nil {
			logger.Error("failed to stop cart task worker: %w", err)
		}
	}()

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
