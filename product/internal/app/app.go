package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/Go-Marketplace/backend/pkg/redis"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/interceptors"
	"github.com/Go-Marketplace/backend/product/internal/infrastructure/repository"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.ProductConfig.Level)

	ctx := context.Background()

	redis, err := redis.New(ctx, cfg.ProductConfig.Redis.Url)
	if err != nil {
		log.Fatalf("failed to run redis.New: %s", err)
	}
	defer redis.Close()

	pg, err := postgres.New(cfg.ProductConfig.PG.URL)
	if err != nil {
		log.Fatalf("failed to run postgres.New: %s", err)
	}
	defer pg.Close()

	// Create cart client
	cartConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.CartConfig.GRPC.Host, cfg.CartConfig.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create cartConn: %v", err)
	}
	defer cartConn.Close()

	cartClient := pbCart.NewCartClient(cartConn)

	discountRepo := repository.NewDiscountRepo(redis, logger)
	productRepo := repository.NewProductRepo(pg, logger)
	productUsecase := usecase.NewProductUsecase(productRepo, discountRepo, logger)
	productHandler := handler.NewProductRoutes(productUsecase, cartClient, logger)

	interceptor := interceptors.NewInterceptorManager(logger)
	grpcServer, err := grpcserver.New(
		cfg.ProductConfig.GRPC.Port,
		grpc.UnaryInterceptor(
			interceptor.LogRequest,
		),
	)
	if err != nil {
		log.Fatalf("failed to create new grcp server: %s", err)
	}
	pbProduct.RegisterProductServer(grpcServer.Server, productHandler)
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
