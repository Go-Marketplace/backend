package interceptors

import (
	"context"
	"strings"
	"time"

	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/gateway/internal/model"
	"github.com/Go-Marketplace/backend/gateway/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type interceptorManager struct {
	rbacManager *model.RBACManager
	jwtManager  *usecase.JWTManager
	logger      *logger.Logger
}

func NewInterceptorManager(
	logger *logger.Logger,
	jwtManager *usecase.JWTManager,
	rbac *model.RBACManager,
) *interceptorManager {
	return &interceptorManager{
		jwtManager:  jwtManager,
		logger:      logger,
		rbacManager: rbac,
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

func (interceptor *interceptorManager) getRequestUserClaim(ctx context.Context) (*controller.UserClaim, error) {
	claim := controller.UserClaim{}

	if token, err := grpc_auth.AuthFromMD(ctx, "bearer"); err == nil {
		payload, err := interceptor.jwtManager.ValidateToken(token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "failed to validate token: %s", err)
		}

		interceptor.logger.Info("payload: %v", payload)

		if err := mapstructure.Decode(payload, &claim); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user claim from payload: %s", err)
		}
	} else {
		errStatus, ok := status.FromError(err)
		if !ok {
			return nil, status.Errorf(codes.Internal, "failed to get status from err: %s", err)
		}

		if errStatus.Code() == codes.Unauthenticated {
			claim.Role = pbUser.UserRole_GUEST
		} else {
			return nil, status.Errorf(codes.InvalidArgument, "failed to get auth from metadata: %s", err)
		}
	}

	return &claim, nil
}

func getRequestMethod(ctx context.Context) (string, error) {
	fullMethod, ok := grpc.Method(ctx)
	if !ok {
		return "", status.Error(codes.Internal, "failed to get grpc method from context")
	}

	methodParts := strings.Split(fullMethod, "/")
	if len(methodParts) != 3 {
		return "", status.Error(codes.Internal, "invalid request scheme")
	}

	return methodParts[2], nil
}

func (interceptor *interceptorManager) AuthRequest(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	claim, err := interceptor.getRequestUserClaim(ctx)
	if err != nil {
		return nil, err
	}

	method, err := getRequestMethod(ctx)
	if err != nil {
		return nil, err
	}

	interceptor.logger.Info("Got claim id: %s, role: %s, method: %s", claim.ID, claim.Role.String(), method)

	if !interceptor.rbacManager.RBAC.IsGranted(claim.Role.String(), interceptor.rbacManager.Permissions[method], nil) {
		return nil, status.Error(codes.PermissionDenied, "Not permited")
	}

	md := metadata.Pairs("user_id", claim.ID)
	md.Set("role", claim.Role.String())
	ctx = metadata.NewOutgoingContext(ctx, md)

	return handler(ctx, req)
}
