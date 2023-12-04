package repository

import (
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func getCartsQuery() sq.SelectBuilder {
	return psql.Select(
		"user_id",
		"created_at",
		"updated_at",
	).
		From("carts")
}

func getFullCartsQuery() sq.SelectBuilder {
	return psql.Select(
		"carts.user_id",
		"carts.created_at",
		"carts.updated_at",
		"cartlines.user_id",
		"cartlines.product_id",
		"cartlines.quantity",
		"cartlines.created_at",
		"cartlines.updated_at",
	).
		From("carts").
		Join("cartlines USING (user_id)")
}

func getFullUserCartQuery(userID uuid.UUID) sq.SelectBuilder {
	return getFullCartsQuery().
		Where(sq.Eq{
			"carts.user_id": userID,
		})
}

func getUserCartQuery(userID uuid.UUID) sq.SelectBuilder {
	return getCartsQuery().
		Where(sq.Eq{
			"carts.user_id": userID,
		})
}

func createCartQuery(cart model.Cart) sq.InsertBuilder {
	return psql.Insert("carts").
		Columns(
			"user_id",
			"created_at",
			"updated_at",
		).
		Values(
			cart.UserID,
			cart.CreatedAt,
			cart.UpdatedAt,
		)
}

func updateCartQuery(userID uuid.UUID) sq.UpdateBuilder {
	return psql.Update("carts").
		Set("updated_at", time.Now()).
		Where(sq.Eq{
			"user_id": userID,
		})
}

func getCartlines() sq.SelectBuilder {
	return psql.Select(
		"user_id",
		"product_id",
		"quantity",
		"created_at",
		"updated_at",
	).
		From("cartlines")
}

func getCartlineQuery(userID uuid.UUID, productID uuid.UUID) sq.SelectBuilder {
	return getCartlines().
		Where(sq.And{
			sq.Eq{
				"user_id": userID,
			},
			sq.Eq{
				"product_id": productID,
			},
		})
}

func createCartlineQuery(cartline *model.CartLine) sq.InsertBuilder {
	return psql.Insert("cartlines").
		Columns(
			"user_id",
			"product_id",
			"quantity",
			"created_at",
			"updated_at",
		).
		Values(
			cartline.UserID,
			cartline.ProductID,
			cartline.Quantity,
			cartline.CreatedAt,
			cartline.UpdatedAt,
		)
}

func updateCartlineQuery(cartline model.CartLine) sq.UpdateBuilder {
	query := psql.Update("cartlines")

	if cartline.Quantity != 0 {
		query = query.Set("quantity", cartline.Quantity)
	}

	return query.
		Set("updated_at", time.Now()).
		Where(sq.And{
			sq.Eq{
				"user_id": cartline.UserID,
			},
			sq.Eq{
				"product_id": cartline.ProductID,
			},
		})
}

func deleteCartQuery(userID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("carts").
		Where(sq.Eq{
			"user_id": userID,
		})
}

func deleteCartlineQuery(userID uuid.UUID, productID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("cartlines").
		Where(sq.And{
			sq.Eq{
				"user_id": userID,
			},
			sq.Eq{
				"product_id": productID,
			},
		})
}

func deleteProductCartlinesQuery(productID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("cartlines").
		Where(sq.Eq{
			"product_id": productID,
		})
}

func deleteCartCartlinesQuery(userID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("cartlines").
		Where(sq.Eq{
			"user_id": userID,
		})
}
