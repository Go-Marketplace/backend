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
		UserId:      resp.ID,
		AccessToken: resp.Token,
	}, nil
}

func (router *gatewayRoutes) Login(ctx context.Context, req *pbGateway.LoginRequest) (*pbGateway.LoginResponse, error) {
	resp, err := controller.Login(ctx, router.userClient, router.jwtManager, req)
	if err != nil {
		return nil, err
	}

	return &pbGateway.LoginResponse{
		UserId:      resp.ID,
		AccessToken: resp.Token,
	}, nil
}

// User

func (router *gatewayRoutes) GetUser(ctx context.Context, req *pbUser.GetUserRequest) (*pbUser.UserResponse, error) {
	return router.userClient.GetUser(ctx, req)
}

func (router *gatewayRoutes) GetUsers(ctx context.Context, req *pbUser.GetUsersRequest) (*pbUser.UsersResponse, error) {
	return router.userClient.GetUsers(ctx, req)
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

func (router *gatewayRoutes) GetOrder(ctx context.Context, req *pbOrder.GetOrderRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.GetOrder(ctx, req)
}

func (router *gatewayRoutes) GetOrders(ctx context.Context, req *pbOrder.GetOrdersRequest) (*pbOrder.OrdersResponse, error) {
	return router.orderClient.GetOrders(ctx, req)
}

func (router *gatewayRoutes) GetUserOrders(ctx context.Context, req *pbGateway.GetUserOrdersRequest) (*pbOrder.OrdersResponse, error) {
	return router.orderClient.GetOrders(ctx, &pbOrder.GetOrdersRequest{
		UserId: req.UserId,
	})
}

func (router *gatewayRoutes) CreateOrder(ctx context.Context, req *pbOrder.CreateOrderRequest) (*pbOrder.OrderResponse, error) {
	return router.orderClient.CreateOrder(ctx, req)
}

func (router *gatewayRoutes) DeleteOrder(ctx context.Context, req *pbOrder.DeleteOrderRequest) (*pbOrder.DeleteOrderResponse, error) {
	return router.orderClient.DeleteOrder(ctx, req)
}

func (router *gatewayRoutes) GetOrderline(ctx context.Context, req *pbOrder.GetOrderlineRequest) (*pbOrder.OrderlineResponse, error) {
	return router.orderClient.GetOrderline(ctx, req)
}

func (router *gatewayRoutes) UpdateOrderline(ctx context.Context, req *pbOrder.UpdateOrderlineRequest) (*pbOrder.OrderlineResponse, error) {
	return router.orderClient.UpdateOrderline(ctx, req)
}

func (router *gatewayRoutes) DeleteOrderline(ctx context.Context, req *pbOrder.DeleteOrderlineRequest) (*pbOrder.DeleteOrderlineResponse, error) {
	return router.orderClient.DeleteOrderline(ctx, req)
}

// Cart

func (router *gatewayRoutes) GetUserCart(ctx context.Context, req *pbCart.GetUserCartRequest) (*pbCart.CartResponse, error) {
	return router.cartClient.GetUserCart(ctx, req)
}

func (router *gatewayRoutes) CreateCartline(ctx context.Context, req *pbCart.CreateCartlineRequest) (*pbCart.CartlineResponse, error) {
	return router.cartClient.CreateCartline(ctx, req)
}

func (router *gatewayRoutes) UpdateCartline(ctx context.Context, req *pbCart.UpdateCartlineRequest) (*pbCart.CartlineResponse, error) {
	return router.cartClient.UpdateCartline(ctx, req)
}

func (router *gatewayRoutes) DeleteCartline(ctx context.Context, req *pbCart.DeleteCartlineRequest) (*pbCart.DeleteCartlineResponse, error) {
	return router.cartClient.DeleteCartline(ctx, req)
}

func (router *gatewayRoutes) DeleteCartCartlines(ctx context.Context, req *pbCart.DeleteCartCartlinesRequest) (*pbCart.DeleteCartCartlinesResponse, error) {
	return router.cartClient.DeleteCartCartlines(ctx, req)
}

// Product

func (router *gatewayRoutes) GetProduct(ctx context.Context, req *pbProduct.GetProductRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.GetProduct(ctx, req)
}

func (router *gatewayRoutes) GetProducts(ctx context.Context, req *pbProduct.GetProductsRequest) (*pbProduct.ProductsResponse, error) {
	return router.productClient.GetProducts(ctx, req)
}

func (router *gatewayRoutes) GetUserProducts(ctx context.Context, req *pbGateway.GetUserProductsRequest) (*pbProduct.ProductsResponse, error) {
	return router.productClient.GetProducts(ctx, &pbProduct.GetProductsRequest{
		UserId: req.UserId,
	})
}

func (router *gatewayRoutes) CreateProduct(ctx context.Context, req *pbProduct.CreateProductRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.CreateProduct(ctx, req)
}

func (router *gatewayRoutes) UpdateProduct(ctx context.Context, req *pbProduct.UpdateProductRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.UpdateProduct(ctx, req)
}

func (router *gatewayRoutes) ModerateProduct(ctx context.Context, req *pbProduct.ModerateProductRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.ModerateProduct(ctx, req)
}

func (router *gatewayRoutes) DeleteProduct(ctx context.Context, req *pbProduct.DeleteProductRequest) (*pbProduct.DeleteProductResponse, error) {
	return router.productClient.DeleteProduct(ctx, req)
}

// Category

func (router *gatewayRoutes) GetCategory(ctx context.Context, req *pbProduct.GetCategoryRequest) (*pbProduct.CategoryResponse, error) {
	return router.productClient.GetCategory(ctx, req)
}

func (router *gatewayRoutes) GetAllCategories(ctx context.Context, req *pbProduct.GetAllCategoriesRequest) (*pbProduct.CategoriesResponse, error) {
	return router.productClient.GetAllCategories(ctx, req)
}

// Discount

func (router *gatewayRoutes) CreateDiscount(ctx context.Context, req *pbProduct.CreateDiscountRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.CreateDiscount(ctx, req)
}

func (router *gatewayRoutes) DeleteDiscount(ctx context.Context, req *pbProduct.DeleteDiscountRequest) (*pbProduct.ProductResponse, error) {
	return router.productClient.DeleteDiscount(ctx, req)
}
