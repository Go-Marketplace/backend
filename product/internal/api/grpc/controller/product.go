package controller

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/product/internal/model"
	"github.com/Go-Marketplace/backend/product/internal/usecase"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func GetAllProducts(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetAllProductsRequest) ([]*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	products, err := productUsecase.GetAllProducts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return products, nil
}

func GetAllUserProducts(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetAllUserProductsRequest) ([]*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user id: %s", err)
	}

	products, err := productUsecase.GetAllUserProducts(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return products, nil
}

func GetAllCategoryProducts(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.GetAllCategoryProductsRequest) ([]*model.Product, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	products, err := productUsecase.GetAllCategoryProducts(ctx, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return products, nil
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
		CategoryID:  req.CategoryId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		UpdatedAt:   time.Now(),
	}

	product, err := productUsecase.UpdateProduct(ctx, newProduct)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err)
	}

	return product, nil
}

func DeleteProduct(ctx context.Context, productUsecase usecase.IProductUsecase, req *pbProduct.DeleteProductRequest) error {
	if req == nil {
		return status.Errorf(codes.InvalidArgument, "Invalid request")
	}

	id, err := uuid.Parse(req.ProductId)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid id: %s", err)
	}

	err = productUsecase.DeleteProduct(ctx, id)
	if err != nil {
		return status.Errorf(codes.Internal, "Internal error: %s", err)
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
		UpdatedAt: time.Now(),
	}

	product, err := productUsecase.ModerateProduct(ctx, newProduct)
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
