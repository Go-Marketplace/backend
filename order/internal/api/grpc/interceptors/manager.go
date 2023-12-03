package interceptors

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type interceptorManager struct {
	logger *logger.Logger
}

func NewInterceptorManager(logger *logger.Logger) *interceptorManager {
	return &interceptorManager{
		logger: logger,
	}
}

func (interceptor *interceptorManager) LogRequest(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	interceptor.logger.Info("Method: %s, Time: %v, Metadata: %v, Err: %v", info.FullMethod, time.Since(start), md, err)

	return reply, err
}
