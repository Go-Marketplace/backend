package repository

import (
	"time"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
	"github.com/Go-Marketplace/backend/order/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func getFullOrdersQuery() sq.SelectBuilder {
	return psql.Select(
		"orders.order_id",
		"orders.user_id",
		"orders.created_at",
		"orders.updated_at",
		"orderlines.order_id",
		"orderlines.product_id",
		"orderlines.name",
		"orderlines.price",
		"orderlines.quantity",
		"orderlines.status",
		"orderlines.created_at",
		"orderlines.updated_at",
	).
		From("orders").
		Join("orderlines USING (order_id)")
}

func getFullOrderQuery(orderID uuid.UUID) sq.SelectBuilder {
	return getFullOrdersQuery().
		Where(sq.Eq{
			"order_id": orderID,
		})
}

func searchOrdersQuery(searchParams dto.SearchOrderDTO) sq.SelectBuilder {
	query := getFullOrdersQuery()

	if searchParams.UserID != uuid.Nil {
		query = query.Where(sq.Eq{
			"user_id": searchParams.UserID,
		})
	}

	return query
}

func createOrderQuery(order *model.Order) sq.InsertBuilder {
	return psql.Insert("orders").
		Columns(
			"order_id",
			"user_id",
			"created_at",
			"updated_at",
		).
		Values(
			order.ID,
			order.UserID,
			order.CreatedAt,
			order.UpdatedAt,
		)
}

func updateOrderQuery(orderID uuid.UUID) sq.UpdateBuilder {
	return psql.Update("orders").
		Set("updated_at", time.Now()).
		Where(sq.Eq{
			"order_id": orderID,
		})
}

func deleteOrderQuery(orderID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("orders").
		Where(sq.Eq{
			"order_id": orderID,
		})
}

func getOrderlinesQuery() sq.SelectBuilder {
	return psql.Select(
		"order_id",
		"product_id",
		"name",
		"price",
		"quantity",
		"status",
		"created_at",
		"updated_at",
	).
		From("orderlines")
}

func getOrderlineQuery(orderID, productID uuid.UUID) sq.SelectBuilder {
	return getOrderlinesQuery().
		Where(sq.And{
			sq.Eq{
				"order_id": orderID,
			},
			sq.Eq{
				"product_id": productID,
			},
		})
}

func createOrderlineQuery(orderline *model.Orderline) sq.InsertBuilder {
	return psql.Insert("orderlines").
		Columns(
			"order_id",
			"product_id",
			"name",
			"price",
			"quantity",
			"status",
			"created_at",
			"updated_at",
		).
		Values(
			orderline.OrderID,
			orderline.ProductID,
			orderline.Name,
			orderline.Price,
			orderline.Quantity,
			orderline.Status,
			orderline.CreatedAt,
			orderline.UpdatedAt,
		)
}

func updateOrderlineQuery(orderline *model.Orderline) sq.UpdateBuilder {
	return psql.Update("orderlines").
		Set("status", orderline.Status).
		Set("updated_at", time.Now()).
		Where(sq.And{
			sq.Eq{
				"order_id": orderline.OrderID,
			},
			sq.Eq{
				"product_id": orderline.ProductID,
			},
		})
}

func deleteOrderlineQuery(orderID, productID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("orderlines").
		Where(sq.And{
			sq.Eq{
				"order_id": orderID,
			},
			sq.Eq{
				"product_id": productID,
			},
		})
}
