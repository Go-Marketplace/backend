package controller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/Go-Marketplace/backend/order/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbOrder "github.com/Go-Marketplace/backend/proto/gen/order"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	pbUser "github.com/Go-Marketplace/backend/proto/gen/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const maxCancellationTime = 24 * time.Hour

func checkOrderAccessPermission(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	orderID uuid.UUID,
) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, fmt.Errorf("failed to get metadata from incoming context")
	}

	roles := md.Get("role")
	if len(roles) == 0 {
		return false, fmt.Errorf("role did not set in metadata")
	}

	if roles[0] != pbUser.UserRole_USER.String() {
		return true, nil
	}

	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return false, fmt.Errorf("user_id did not set in metadata")
	}

	order, err := orderUsecase.GetOrder(ctx, orderID)
	if err != nil {
		return false, fmt.Errorf("failed to get %v order", orderID)
	}

	if order == nil {
		return false, fmt.Errorf("order not found")
	}

	log.Printf("Check Access Permissions for %s user\n", userIDs[0])

	return userIDs[0] == order.UserID.String(), nil
}

func checkOrderlineAccessPermission(
	ctx context.Context,
	productClient pbProduct.ProductClient,
	productID uuid.UUID,
) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, fmt.Errorf("failed to get metadata from incoming context")
	}

	roles := md.Get("role")
	if len(roles) == 0 {
		return false, fmt.Errorf("role did not set in metadata")
	}

	if roles[0] != pbUser.UserRole_USER.String() {
		return true, nil
	}

	userIDs := md.Get("user_id")
	if len(userIDs) == 0 {
		return false, fmt.Errorf("user_id did not set in metadata")
	}

	md = metadata.Pairs("role", pbUser.UserRole_ADMIN.String())
	ctx = metadata.NewOutgoingContext(ctx, md)
	productResp, err := productClient.GetProduct(ctx, &pbProduct.GetProductRequest{
		ProductId: productID.String(),
	})
	if err != nil {
		return false, fmt.Errorf("failed to get product: %s", err)
	}

	log.Printf("Check Access Permissions for %s user\n", userIDs[0])

	return userIDs[0] == productResp.UserId, nil
}

func GetOrder(ctx context.Context, orderUsecase usecase.IOrderUsecase, req *pbOrder.GetOrderRequest) (*model.Order, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	allow, err := checkOrderAccessPermission(ctx, orderUsecase, orderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check order access permission: %s", err)
	}

	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Access denied")
	}

	order, err := orderUsecase.GetOrder(ctx, orderID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get order: %s", err)
	}

	if order == nil {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return order, nil
}

func GetOrders(ctx context.Context, orderUsecase usecase.IOrderUsecase, req *pbOrder.GetOrdersRequest) ([]*model.Order, error) {
	var err error
	var userID uuid.UUID
	if req.UserId != "" {
		userID, err = uuid.Parse(req.UserId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
		}
	}

	searchParams := dto.SearchOrderDTO{
		UserID: userID,
	}

	orders, err := orderUsecase.GetOrders(ctx, searchParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get order: %s", err)
	}

	return orders, nil
}

func getProducts(ctx context.Context, productClient pbProduct.ProductClient, productIDs []string) ([]*pbProduct.ProductResponse, error) {
	products := make([]*pbProduct.ProductResponse, len(productIDs))
	for i, productID := range productIDs {
		md := metadata.Pairs("role", pbUser.UserRole_ADMIN.String())
		ctx = metadata.NewOutgoingContext(ctx, md)
		productResp, err := productClient.GetProduct(ctx, &pbProduct.GetProductRequest{
			ProductId: productID,
		})
		if err != nil {
			statusErr, ok := status.FromError(err)
			if !ok {
				return nil, fmt.Errorf("failed to get status from err: %w", err)
			}

			if statusErr.Code() != codes.NotFound {
				return nil, fmt.Errorf("failed to get product %s: %w", productID, err)
			}
		}

		products[i] = productResp
	}

	return products, nil
}

func CreateOrder(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	cartClient pbCart.CartClient,
	productClient pbProduct.ProductClient,
	req *pbOrder.CreateOrderRequest,
) (*model.Order, error) {
	md := metadata.Pairs("role", pbUser.UserRole_ADMIN.String())
	ctx = metadata.NewOutgoingContext(ctx, md)
	cartResp, err := cartClient.GetUserCart(ctx, &pbCart.GetUserCartRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get cart: %s", err)
	}

	productIDs := make([]string, 0, len(cartResp.Cartlines))
	for _, cartline := range cartResp.Cartlines {
		productIDs = append(productIDs, cartline.ProductId)
	}

	products, err := getProducts(ctx, productClient, productIDs)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get products: %s", err)
	}

	userID, err := uuid.Parse(cartResp.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	newOrder := &model.Order{
		ID:         uuid.New(),
		UserID:     userID,
		Orderlines: make([]*model.Orderline, 0, len(cartResp.Cartlines)),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	for i, cartline := range cartResp.Cartlines {
		if products[i] == nil {
			continue
		}

		productID, err := uuid.Parse(cartline.ProductId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
		}

		orderline := &model.Orderline{
			OrderID:   newOrder.ID,
			ProductID: productID,
			Name:      products[i].Name,
			Quantity:  cartline.Quantity,
			Price:     products[i].Price,
			Status:    model.PendingPayment,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		newOrder.Orderlines = append(newOrder.Orderlines, orderline)
	}

	order, err := orderUsecase.CreateOrder(ctx, newOrder)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create order: %s", err)
	}

	if _, err = cartClient.PrepareOrder(ctx, &pbCart.PrepareOrderRequest{
		UserId: cartResp.UserId,
	}); err != nil {
		if errDelete := orderUsecase.DeleteOrder(ctx, order.ID); errDelete != nil {
			return nil, status.Errorf(codes.Internal, "Failed to delete order: %s", errDelete)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete cart cartlines: %s", err)
	}

	return order, nil
}

func returnProducts(ctx context.Context, productClient pbProduct.ProductClient, orderlines ...*model.Orderline) error {
	productIDs := make([]string, 0, len(orderlines))
	for _, orderline := range orderlines {
		productIDs = append(productIDs, orderline.ProductID.String())
	}

	products, err := getProducts(ctx, productClient, productIDs)
	if err != nil {
		return fmt.Errorf("failed to get products: %w", err)
	}

	updateProductRequests := make([]*pbProduct.UpdateProductRequest, 0, len(orderlines))
	for i, orderline := range orderlines {
		if products[i] != nil {
			newQuantity := products[i].Quantity + orderline.Quantity
			updateProductRequests = append(updateProductRequests, &pbProduct.UpdateProductRequest{
				ProductId: orderline.ProductID.String(),
				Quantity:  &newQuantity,
			})
		}
	}

	md := metadata.Pairs("role", pbUser.UserRole_ADMIN.String())
	ctx = metadata.NewOutgoingContext(ctx, md)
	if _, err := productClient.UpdateProducts(ctx, &pbProduct.UpdateProductsRequest{
		Products: updateProductRequests,
	}); err != nil {
		return fmt.Errorf("failed to update products: %w", err)
	}

	return nil
}

func DeleteOrder(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	productClient pbProduct.ProductClient,
	req *pbOrder.DeleteOrderRequest,
) error {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	allow, err := checkOrderAccessPermission(ctx, orderUsecase, orderID)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check order access permission: %s", err)
	}

	if !allow {
		return status.Errorf(codes.PermissionDenied, "Access denied")
	}

	order, err := orderUsecase.GetOrder(ctx, orderID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get order: %s", err)
	}

	if order == nil {
		return status.Errorf(codes.NotFound, "Order not found")
	}

	if time.Since(order.CreatedAt) > maxCancellationTime {
		return status.Errorf(codes.Canceled, "Maximum order cancellation time has been exceeded")
	}

	if err = orderUsecase.DeleteOrder(ctx, orderID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete order: %s", err)
	}

	if err = returnProducts(ctx, productClient, order.Orderlines...); err != nil {
		if _, errCreate := orderUsecase.CreateOrder(ctx, order); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to recreate order: %s", errCreate)
		}
		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}

func DeleteUserOrders(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	productClient pbProduct.ProductClient,
	req *pbOrder.DeleteUserOrdersRequest,
) error {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	orders, err := orderUsecase.GetOrders(ctx, dto.SearchOrderDTO{
		UserID: userID,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get user orders: %s", err)
	}

	orderlines := make([]*model.Orderline, 0)
	for _, order := range orders {
		orderlines = append(orderlines, order.Orderlines...)
	}

	if err = orderUsecase.DeleteUserOrders(ctx, userID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete user orders: %s", err)
	}

	if err = returnProducts(ctx, productClient, orderlines...); err != nil {
		for _, order := range orders {
			if _, errCreate := orderUsecase.CreateOrder(ctx, order); errCreate != nil {
				return status.Errorf(codes.Internal, "Failed to recreate order: %s", errCreate)
			}
		}

		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}

func GetOrderline(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	productClient pbProduct.ProductClient,
	req *pbOrder.GetOrderlineRequest,
) (*model.Orderline, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	allow, err := checkOrderlineAccessPermission(ctx, productClient, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check orderline access permission: %s", err)
	}

	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Access denied")
	}

	orderline, err := orderUsecase.GetOrderline(ctx, orderID, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get orderline: %s", err)
	}

	if orderline == nil {
		return nil, status.Errorf(codes.NotFound, "Orderline not found")
	}

	return orderline, nil
}

func UpdateOrderline(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	productClient pbProduct.ProductClient,
	req *pbOrder.UpdateOrderlineRequest,
) (*model.Orderline, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	allow, err := checkOrderlineAccessPermission(ctx, productClient, productID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check orderline access permission: %s", err)
	}

	if !allow {
		return nil, status.Errorf(codes.PermissionDenied, "Access denied")
	}

	newOrderline := &model.Orderline{
		OrderID:   orderID,
		ProductID: productID,
		Status:    model.OrderlineStatus(req.Status),
	}

	orderline, err := orderUsecase.UpdateOrderline(ctx, newOrderline)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update orderline: %s", err)
	}

	if orderline == nil {
		return nil, status.Errorf(codes.NotFound, "Orderline not found")
	}

	return orderline, nil
}

func DeleteOrderline(
	ctx context.Context,
	orderUsecase usecase.IOrderUsecase,
	productClient pbProduct.ProductClient,
	req *pbOrder.DeleteOrderlineRequest,
) error {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid order id: %s", err)
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	allow, err := checkOrderlineAccessPermission(ctx, productClient, productID)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to check orderline access permission: %s", err)
	}

	if !allow {
		return status.Errorf(codes.PermissionDenied, "Access denied")
	}

	orderline, err := orderUsecase.GetOrderline(ctx, orderID, productID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get orderline: %s", err)
	}

	if orderline == nil {
		return status.Errorf(codes.NotFound, "Orderline not found")
	}

	if err = orderUsecase.DeleteOrderline(ctx, orderID, productID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete orderline: %s", err)
	}

	if err = returnProducts(ctx, productClient, orderline); err != nil {
		if errCreate := orderUsecase.CreateOrderline(ctx, orderline); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to create orderline: %s", errCreate)
		}
		return status.Errorf(codes.Internal, "Failed to return products: %s", err)
	}

	return nil
}
