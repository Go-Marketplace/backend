package app

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/handler"
	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/interceptors"
	"github.com/Go-Marketplace/backend/gateway/internal/model"
	"github.com/Go-Marketplace/backend/gateway/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/grpcserver"
	"github.com/Go-Marketplace/backend/pkg/httpserver"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mikespook/gorbac"
	"github.com/xiam/to"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

//go:embed gateway.swagger.json
var spec []byte

func getRBAC(rolesPath string, inheritancePath string) (*model.RBACManager, error) {
	rolesFile, err := os.Open(rolesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file: %w", rolesPath, err)
	}
	defer rolesFile.Close()

	inheritanceFile, err := os.Open(inheritancePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file: %w", inheritancePath, err)
	}
	defer inheritanceFile.Close()

	roles := make(map[string][]string)
	inheritance := make(map[string][]string)

	if err := json.NewDecoder(rolesFile).Decode(&roles); err != nil {
		return nil, fmt.Errorf("failed to decode from %s file", rolesPath)
	}

	if err := json.NewDecoder(inheritanceFile).Decode(&inheritance); err != nil {
		return nil, fmt.Errorf("failed to decode from %s file", inheritancePath)
	}

	rbac := gorbac.New()

	permissions := make(gorbac.Permissions)

	for roleID, permissionIDs := range roles {
		role := gorbac.NewStdRole(roleID)
		for _, permissionID := range permissionIDs {
			if _, ok := permissions[permissionID]; !ok {
				permissions[permissionID] = gorbac.NewStdPermission(permissionID)
			}
			if err = role.Assign(permissions[permissionID]); err != nil {
				return nil, fmt.Errorf("failed to assign role permission: %s", permissionID)
			}
		}
		if err = rbac.Add(role); err != nil {
			return nil, fmt.Errorf("failed to add role: %s", role.ID())
		}
	}

	for roleID, parentIDs := range inheritance {
		if err := rbac.SetParents(roleID, parentIDs); err != nil {
			return nil, fmt.Errorf("failed to set parents for %s", roleID)
		}
	}

	return model.NewRBACManager(rbac, permissions), nil
}

func Run(cfg *config.Config) {
	logger := logger.New(cfg.GatewayConfig.Level)

	jwtManager, err := usecase.NewJWTManager(
		cfg.GatewayConfig.Auth.AccessTokenPrivateKey,
		cfg.GatewayConfig.Auth.AccessTokenPublicKey,
		to.Duration(cfg.GatewayConfig.Auth.AccessTokenMaxAge),
	)
	if err != nil {
		log.Fatalf("failed to create jwtManager: %s", err)
	}

	// Create order client
	orderConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.OrderConfig.GRPC.Host, cfg.OrderConfig.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create orderConn: %s", err)
	}
	defer orderConn.Close()

	orderClient := pbOrder.NewOrderClient(orderConn)

	// Create user client
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.UserConfig.GRPC.Host, cfg.UserConfig.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create userConn: %v", err)
	}
	defer userConn.Close()

	userClient := pbUser.NewUserClient(userConn)

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

	// Create product client
	productConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.ProductConfig.GRPC.Host, cfg.ProductConfig.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to create productConn: %v", err)
	}
	defer productConn.Close()

	productClient := pbProduct.NewProductClient(productConn)

	// Create gateway handler

	gatewayHandler := handler.NewGatewayRoutes(
		orderClient,
		userClient,
		cartClient,
		productClient,
		jwtManager,
		logger,
	)

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %s", err)
	}

	rolesPath := filepath.Join(curDir, "config", "rbac", "roles.json")
	inheritancePath := filepath.Join(curDir, "config", "rbac", "inheritance.json")

	rbacManager, err := getRBAC(rolesPath, inheritancePath)
	if err != nil {
		log.Fatalf("failed to get RBAC: %s", err)
	}

	interceptor := interceptors.NewInterceptorManager(logger, jwtManager, rbacManager)

	// Start GRPC Server
	grpcServer, err := grpcserver.New(
		cfg.GatewayConfig.GRPC.Port,
		grpc.ChainUnaryInterceptor(
			interceptor.LogRequest,
			interceptor.AuthRequest,
		),
	)
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
	httpMux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	err = pbGateway.RegisterGatewayHandlerFromEndpoint(
		context.Background(),
		gwmux,
		fmt.Sprintf("localhost:%v", cfg.GatewayConfig.GRPC.Port),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to register api gateway handler from endpoint: %s", err)
	}

	httpMux.Handle("/", gwmux)
	httpMux.Handle("/api/v1/swagger/", http.StripPrefix("/api/v1/swagger", swaggerui.Handler(spec)))

	httpServer := httpserver.New(httpMux)
	httpServer.Start()
	defer func() {
		if err = httpServer.Shutdown(); err != nil {
			logger.Error("failed to shutdown http server: %s", err)
		}
	}()
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
