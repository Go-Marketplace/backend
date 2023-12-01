package repository

import (
	"time"

	"github.com/Go-Marketplace/backend/product/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/product/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func getAllProductsQuery() sq.SelectBuilder {
	return psql.
		Select(
			"product_id",
			"user_id",
			"category_id",
			"name",
			"description",
			"price",
			"quantity",
			"moderated",
			"created_at",
			"updated_at",
		).
		From("products")
}

func getProductQuery(productID uuid.UUID) sq.SelectBuilder {
	return getAllProductsQuery().
		Where(sq.Eq{
			"product_id": productID,
		})
}

func getSearchProductsQuery(searchParams dto.SearchProductsDTO) sq.SelectBuilder {
	query := getAllProductsQuery().Where(sq.Eq{
		"moderated": searchParams.Moderated,
	})

	if searchParams.UserID != uuid.Nil {
		query = query.Where(sq.Eq{
			"user_id": searchParams.UserID,
		})
	}

	if searchParams.CategoryID != 0 {
		query = query.Where(sq.Eq{
			"category_id": searchParams.CategoryID,
		})
	}

	return query
}

func createProductQuery(product model.Product) sq.InsertBuilder {
	return psql.
		Insert("products").
		Columns(
			"product_id",
			"user_id",
			"category_id",
			"name",
			"description",
			"price",
			"quantity",
			"moderated",
			"created_at",
			"updated_at",
		).
		Values(
			product.ID,
			product.UserID,
			product.CategoryID,
			product.Name,
			product.Description,
			product.Price,
			product.Quantity,
			product.Moderated,
			product.CreatedAt,
			product.UpdatedAt,
		)
}

func updateProductQuery(product model.Product) sq.UpdateBuilder {
	query := psql.Update("products")

	if product.CategoryID != 0 {
		query = query.Set("category_id", product.CategoryID)
	}

	if product.Name != "" {
		query = query.Set("name", product.Name)
	}

	if product.Description != "" {
		query = query.Set("description", product.Description)
	}

	if product.Price != 0 {
		query = query.Set("price", product.Price)
	}

	if product.Quantity != 0 {
		query = query.Set("quantity", product.Quantity)
	}

	if product.Moderated {
		query = query.Set("moderated", product.Moderated)
	}

	query.Set("updated_at", time.Now())

	query = query.Where(sq.Eq{
		"product_id": product.ID,
	})

	return query
}

func deleteProductQuery(productID uuid.UUID) sq.DeleteBuilder {
	return psql.
		Delete("products").
		Where(sq.Eq{
			"product_id": productID,
		})
}

func getAllCategoriesQuery() sq.SelectBuilder {
	return psql.
		Select(
			"category_id",
			"name",
			"description",
		).
		From("categories")
}

func getCategory(categoryID int32) sq.SelectBuilder {
	return getAllCategoriesQuery().
		Where(sq.Eq{
			"category_id": categoryID,
		})
}
