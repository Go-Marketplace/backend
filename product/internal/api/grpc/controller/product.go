package controller

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/product/internal"
	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	pbCart "github.com/Go-Marketplace/backend/proto/gen/cart"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetProducts(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetProductsRequest) ([]*model.Product, error) {
	var err error
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	var userID uuid.UUID
	if req.UserId != "" {
		userID, err = uuid.Parse(req.UserId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
		}
	}

	searchParams := dto.SearchProductsDTO{
		UserID:     userID,
		CategoryID: req.CategoryId,
		Moderated:  req.Moderated,
	}

	products, err := productUsecase.GetProducts(ctx, searchParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return products, nil
}

func GetProduct(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetProductRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	id, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	product, err := productUsecase.GetProduct(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if product == nil {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	return product, nil
}

func CreateProduct(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.CreateProductRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	product := model.Product{
		ID:          uuid.New(),
		UserID:      userID,
		CategoryID:  req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Moderated:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = productUsecase.CreateProduct(ctx, product)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return &product, nil
}

func UpdateProduct(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.UpdateProductRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	newProduct := model.Product{
		ID:          productID,
		CategoryID:  internal.Unwrap(req.CategoryId),
		Name:        internal.Unwrap(req.Name),
		Description: internal.Unwrap(req.Description),
		Price:       internal.Unwrap(req.Price),
		Quantity:    internal.Unwrap(req.Quantity),
	}

	product, err := productUsecase.UpdateProduct(ctx, newProduct)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return product, nil
}

func UpdateProducts(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.UpdateProductsRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	products := make([]model.Product, 0, len(req.Products))

	for _, reqProduct := range req.Products {
		productID, err := uuid.Parse(reqProduct.ProductId)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
		}

		products = append(products, model.Product{
			ID:          productID,
			CategoryID:  internal.Unwrap(reqProduct.CategoryId),
			Name:        internal.Unwrap(reqProduct.Name),
			Description: internal.Unwrap(reqProduct.Description),
			Price:       internal.Unwrap(reqProduct.Price),
			Quantity:    internal.Unwrap(reqProduct.Quantity),
		})
	}

	if err := productUsecase.UpdateProducts(ctx, products); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return nil
}

func DeleteProduct(
	ctx context.Context,
	productUsecase usecase.IProductUsecase,
	cartClient pbCart.CartClient,
	req *pbProduct.DeleteProductRequest,
) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	product, err := productUsecase.GetProduct(ctx, productID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get product: %s", err)
	}

	if product == nil {
		return status.Errorf(codes.NotFound, "Product not found")
	}

	if err = productUsecase.DeleteProduct(ctx, productID); err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if _, err = cartClient.DeleteProductCartlines(ctx, &pbCart.DeleteProductCartlinesRequest{
		ProductId: req.ProductId,
	}); err != nil {
		if errCreate := productUsecase.CreateProduct(ctx, *product); errCreate != nil {
			return status.Errorf(codes.Internal, "Failed to create product: %s", err)
		}
		return status.Errorf(codes.Internal, "Failed to delete product cartlines: %s", err)
	}

	return nil
}

func DeleteUserProducts(
	ctx context.Context,
	productUsecase usecase.IProductUsecase,
	cartClient pbCart.CartClient,
	req *pbProduct.DeleteUserProductsRequest,
) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	products, err := productUsecase.GetProducts(ctx, dto.SearchProductsDTO{
		UserID: userID,
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get products: %s", err)
	}

	if err = productUsecase.DeleteUserProducts(ctx, userID); err != nil {
		return status.Errorf(codes.Internal, "Failed to delete user products: %s", err)
	}

	for _, product := range products {
		if _, err = cartClient.DeleteProductCartlines(ctx, &pbCart.DeleteProductCartlinesRequest{
			ProductId: product.ID.String(),
		}); err != nil {
			return status.Errorf(codes.Internal, "Failed to delete product cartlines: %s", err)
		}
	}

	return nil
}

func ModerateProduct(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.ModerateProductRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	newProduct := model.Product{
		ID:        productID,
		Moderated: req.Moderated,
	}

	product, err := productUsecase.UpdateProduct(ctx, newProduct)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if product == nil {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	return product, nil
}

func GetCategory(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetCategoryRequest) (*model.Category, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	category, err := productUsecase.GetCategory(ctx, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if category == nil {
		return nil, status.Errorf(codes.NotFound, "Category not found")
	}

	return category, nil
}

func GetAllCategories(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetAllCategoriesRequest) ([]*model.Category, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	categories, err := productUsecase.GetAllCategories(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return categories, nil
}

func CreateDiscount(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.CreateDiscountRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid product id: %s", err)
	}

	discount := model.Discount{
		ProductID: productID,
		Percent:   req.Percent,
		CreatedAt: time.Now(),
		EndedAt:   req.EndedAt.AsTime(),
	}

	product, err := productUsecase.CreateDiscount(ctx, discount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return product, nil
}

func DeleteDiscount(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.DeleteDiscountRequest) (*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	product, err := productUsecase.DeleteDiscount(ctx, productID)
	if err != nil {
		status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	if product == nil {
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	return product, nil
}
