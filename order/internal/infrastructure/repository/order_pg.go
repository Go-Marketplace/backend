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

func scanFullOrder(rows pgx.Rows, order *model.Order, orderline *model.Orderline) error {
	return rows.Scan(
		&order.ID,
		&order.UserID,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
		&orderline.ID,
		&orderline.OrderID,
		&orderline.ProductID,
		&orderline.Name,
		&orderline.Price,
		&orderline.Quantity,
		&orderline.Status,
		&orderline.CreatedAt,
		&orderline.UpdatedAt,
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
		orderline := &model.Orderline{}

		err := scanFullOrder(rows, order, orderline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full order: %w", err)
		}

		if oldOrder, ok := orderMap[order.ID.String()]; ok {
			order = oldOrder
		} else {
			order.Orderlines = make([]*model.Orderline, 0)
		}

		order.Orderlines = append(order.Orderlines, orderline)
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
		return nil, fmt.Errorf("failed to Query getAllUserOrders: %w", err)
	}
	defer rows.Close()

	ordersMap := make(map[string]*model.Order)
	for rows.Next() {
		order := &model.Order{}
		orderline := &model.Orderline{}

		err := scanFullOrder(rows, order, orderline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full order: %w", err)
		}

		if oldOrder, ok := ordersMap[order.ID.String()]; ok {
			order = oldOrder
		} else {
			order.Orderlines = make([]*model.Orderline, 0)
		}

		order.Orderlines = append(order.Orderlines, orderline)
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
		order.TotalPrice,
		order.CreatedAt,
		order.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createOrder: %w", err)
	}

	for _, orderline := range order.Orderlines {
		_, err := repo.pg.Pool.Exec(
			ctx,
			createOrderline,
			orderline.ID,
			orderline.OrderID,
			orderline.ProductID,
			orderline.Name,
			orderline.Price,
			orderline.Quantity,
			orderline.Status,
			orderline.CreatedAt,
			orderline.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("failed to Exec createOrderline: %w", err)
		}
	}

	return nil
}

func (repo *OrderRepo) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteOrder, id)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteOrder: %w", err)
	}

	return nil
}
