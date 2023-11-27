package handler

import (
	"context"

	"github.com/Go-Marketplace/backend/gateway/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/gateway/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbGateway "github.com/Go-Marketplace/backend/proto/gen/gateway"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
)

type gatewayRoutes struct {
	pbGateway.GatewayServer

	orderClient   pbOrder.OrderClient
	userClient    pbUser.UserClient
	cartClient    pbCart.CartClient
	productClient pbProduct.ProductClient
	jwtManager    *usecase.JWTManager
	logger        *logger.Logger
}

func NewGatewayRoutes(
	orderClient pbOrder.OrderClient,
	userClient pbUser.UserClient,
	cartClient pbCart.CartClient,
	productClient pbProduct.ProductClient,
	jwtManager *usecase.JWTManager,
	logger *logger.Logger,
) *gatewayRoutes {
	return &gatewayRoutes{
		orderClient:   orderClient,
		userClient:    userClient,
		cartClient:    cartClient,
		productClient: productClient,
		jwtManager:    jwtManager,
		logger:        logger,
	}
}

// Auth

func (router *gatewayRoutes) RegisterUser(ctx context.Context, req *pbGateway.RegisterUserRequest) (*pbGateway.RegisterUserResponse, error) {
	resp, err := controller.RegisterUser(ctx, router.userClient, router.jwtManager, req)
	if err != nil {
		return nil, err
	}

	return &pbGateway.RegisterUserResponse{
		UserId:      resp.ID.String(),
		AccessToken: resp.Token,
	}, nil
}

func (router *gatewayRoutes) Login(ctx context.Context, req *pbGateway.LoginRequest) (*pbGateway.LoginResponse, error) {
	token, err := controller.Login(ctx, router.userClient, router.jwtManager, req)
	if err != nil {
		return nil, err
	}

	return &pbGateway.LoginResponse{
		AccessToken: token,
	}, nil
}

// User

func (router *gatewayRoutes) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.UserResponse, error) {
	return router.userClient.GetUser(ctx, req)
}

func (router *gatewayRoutes) GetAllUsers(ctx context.Context, req *pbUser.GetAllUsersRequest) (*pbUser.GetAllUsersResponse, error) {
	return router.userClient.GetAllUsers(ctx, req)
}

func (router *gatewayRoutes) UpdateUser(ctx context.Context, req *pbUser.UpdateUserRequest) (*pbUser.UserResponse, error) {
	return router.userClient.UpdateUser(ctx, req)
}

func (router *gatewayRoutes) ChangeUserRole(ctx context.Context, req *pbUser.ChangeUserRoleRequest) (*pbUser.UserResponse, error) {
	return router.userClient.ChangeUserRole(ctx, req)
}

func (router *gatewayRoutes) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*pbUser.DeleteUserResponse, error) {
	return router.userClient.DeleteUser(ctx, req)
}

// Order

func (router *gatewayRoutes) CreateOrder(ctx context.Context, req *pbOrder.CreateOrderRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.CreateOrder(ctx, req)
}

func (router *gatewayRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.GetOrder(ctx, req)
}

func (router *gatewayRoutes) GetAllUserOrders(ctx context.Context, req *pbOrder.GetAllUserOrdersRequest) (*pbOrder.GetAllUserOrdersResponse, error) {
	return router.orderClient.GetAllUserOrders(ctx, req)
}

func (router *gatewayRoutes) DeleteOrder(ctx context.Context, req *pbOrder.DeleteOrderRequest) (*pbOrder.DeleteOrderResponse, error) {
	return router.orderClient.DeleteOrder(ctx, req)
}

func (router *gatewayRoutes) UpdateOrderline(ctx context.Context, req *pbOrder.UpdateOrderlineRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.UpdateOrderline(ctx, req)
}

// Cart

func (router *gatewayRoutes) GetCart(ctx context.Context, req *pbCart.GetCartRequest) (*pbCart.CartModel, error) {
	return router.cartClient.GetCart(ctx, req)
}

func (router *gatewayRoutes) GetUserCart(ctx context.Context, req *pbCart.GetUserCartRequest) (*pbCart.CartModel, error) {
	return router.cartClient.GetUserCart(ctx, req)
}

func (router *gatewayRoutes) CreateCartline(ctx context.Context, req *pbCart.CreateCartlineRequest) (*pbCart.CreateCartlineResponse, error) {
	return router.cartClient.CreateCartline(ctx, req)
}

func (router *gatewayRoutes) UpdateCartline(ctx context.Context, req *pbCart.UpdateCartlineRequest) (*pbCart.CartModel, error) {
	return router.cartClient.UpdateCartline(ctx, req)
}

func (router *gatewayRoutes) DeleteCartline(ctx context.Context, req *pbCart.DeleteCartlineRequest) (*pbCart.DeleteCartlineResponse, error) {
	return router.cartClient.DeleteCartline(ctx, req)
}

// Product

func (router *gatewayRoutes) GetProduct(ctx context.Context, req *pbProduct.GetProductRequest) (*pbProduct.ProductModel, error) {
	return router.productClient.GetProduct(ctx, req)
}

func (router *gatewayRoutes) GetAllProduct(ctx context.Context, req *pbProduct.GetAllProductsRequest) (*pbProduct.GetProductsResponse, error) {
	return router.productClient.GetAllProducts(ctx, req)
}

func (router *gatewayRoutes) GetAllUserProducts(ctx context.Context, req *pbProduct.GetAllUserProductsRequest) (*pbProduct.GetProductsResponse, error) {
	return router.productClient.GetAllUserProducts(ctx, req)
}

func (router *gatewayRoutes) CreateProduct(ctx context.Context, req *pbProduct.CreateProductRequest) (*pbProduct.ProductModel, error) {
	return router.productClient.CreateProduct(ctx, req)
}

func (router *gatewayRoutes) UpdateProduct(ctx context.Context, req *pbProduct.UpdateProductRequest) (*pbProduct.ProductModel, error) {
	return router.productClient.UpdateProduct(ctx, req)
}

func (router *gatewayRoutes) ModerateProduct(ctx context.Context, req *pbProduct.ModerateProductRequest) (*pbProduct.ProductModel, error) {
	return router.productClient.ModerateProduct(ctx, req)
}

func (router *gatewayRoutes) DeleteProduct(ctx context.Context, req *pbProduct.DeleteProductRequest) (*pbProduct.DeleteProductResponse, error) {
	return router.productClient.DeleteProduct(ctx, req)
}

// Category

func (router *gatewayRoutes) GetCategory(ctx context.Context, req *pbProduct.GetCategoryRequest) (*pbProduct.CategoryModel, error) {
	return router.productClient.GetCategory(ctx, req)
}

func (router *gatewayRoutes) GetAllCategories(ctx context.Context, req *pbProduct.GetAllCategoriesRequest) (*pbProduct.GetAllCategoriesResponse, error) {
	return router.productClient.GetAllCategories(ctx, req)
}

// Discount

func (router *gatewayRoutes) CreateDiscount(ctx context.Context, req *pbProduct.CreateDiscountRequest) (*pbProduct.ProductModel, error) {
	return router.productClient.CreateDiscount(ctx, req)
}

func (router *gatewayRoutes) DeleteDiscount(ctx context.Context, req *pbProduct.DeleteDiscountRequest) (*pbProduct.DeleteDiscountResponse, error) {
	return router.productClient.DeleteDiscount(ctx, req)
}
