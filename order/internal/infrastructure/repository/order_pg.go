package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/order/internal/model"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OrderRepo struct {
	pg     *postgres.Postgres
	logger *logger.Logger
}

func NewOrderRepo(pg *postgres.Postgres, logger *logger.Logger) *OrderRepo {
	return &OrderRepo{
		pg:     pg,
		logger: logger,
	}
}

func scanFullOrder(rows pgx.Rows, order *model.Order, cartline *model.Cartline, product *model.Product) error {
	return rows.Scan(
		&order.ID,
		&order.UserID,
		&order.Status,
		&order.TotalPrice,
		&order.ShippingCost,
		&order.DeliveryAddress,
		&order.DeliveryType,
		&order.CreatedAt,
		&order.UpdatedAt,
		&cartline.ID,
		&cartline.Quantity,
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
	)
}

func (repo *OrderRepo) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to Acquire in GetOrder: %w", err)
    }
    defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetOrder transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				repo.logger.Error("failed to rollback transaction", err)
			}
		} else {
			if err := tx.Commit(ctx); err != nil {
				repo.logger.Error("failed to commit transaction", err)
			}
		}
	}()

	rows, err := tx.Query(ctx, getFullOrderByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullOrderByID: %w", err)
	}
	defer rows.Close()

	orderMap := make(map[string]*model.Order)
	for rows.Next() {
		order := &model.Order{}
		cartline := &model.Cartline{}
		product := &model.Product{}

		err := scanFullOrder(rows, order, cartline, product)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full order: %w", err)
		}

		if oldOrder, ok := orderMap[order.ID.String()]; ok {
			order = oldOrder
		} else {
			order.Cartlines = make([]*model.Cartline, 0)
		}

		cartline.Product = product
		order.Cartlines = append(order.Cartlines, cartline)
		orderMap[order.ID.String()] = order
	}

	return orderMap[id.String()], nil
}

func (repo *OrderRepo) GetAllUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to Acquire in GetOrder: %w", err)
    }
    defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetOrder transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(ctx); err != nil {
				repo.logger.Error("failed to rollback transaction", err)
			}
		} else {
			if err := tx.Commit(ctx); err != nil {
				repo.logger.Error("failed to commit transaction", err)
			}
		}
	}()

	rows, err := tx.Query(ctx, getAllUserOrders, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullOrderByID: %w", err)
	}
	defer rows.Close()

	ordersMap := make(map[string]*model.Order)
	for rows.Next() {
		order := &model.Order{}
		cartline := &model.Cartline{}
		product := &model.Product{}

		err := scanFullOrder(rows, order, cartline, product)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full order: %w", err)
		}

		if oldOrder, ok := ordersMap[order.ID.String()]; ok {
			order = oldOrder
		} else {
			order.Cartlines = make([]*model.Cartline, 0)
		}

		cartline.Product = product
		order.Cartlines = append(order.Cartlines, cartline)
		ordersMap[order.ID.String()] = order
	}

	userOrders := make([]*model.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		userOrders = append(userOrders, order)
	}

	return userOrders, nil
}

func (repo *OrderRepo) CreateOrder(ctx context.Context, order model.Order) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		createOrder,
		order.ID,
		order.UserID,
		order.Status,
		order.TotalPrice,
		order.ShippingCost,
		order.DeliveryAddress,
		order.DeliveryType,
		order.CreatedAt,
		order.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createOrder: %w", err)
	}

	for _, cartline := range order.Cartlines {
		_, err := repo.pg.Pool.Exec(
			ctx,
			createCartline,
			cartline.ID,
			cartline.OrderID,
			cartline.Quantity,
		)
		if err != nil {
			return fmt.Errorf("failed to Exec createCartline: %w", err)
		}

		_, err = repo.pg.Pool.Exec(
			ctx,
			createProduct,
			cartline.Product.ID,
			cartline.Product.CartlineID,
			cartline.Product.Name,
			cartline.Product.Description,
			cartline.Product.Price,
		)
		if err != nil {
			return fmt.Errorf("failed to Exec createProduct: %w", err)
		}
	}

	return nil
}

func (repo *OrderRepo) CancelOrder(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, cancelOrder, id)
	if err != nil {
		return fmt.Errorf("failed to Exec cancelOrder: %w", err)
	}

	return nil
}
