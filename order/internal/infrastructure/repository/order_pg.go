package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/order/internal/api/grpc/dto"
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
		&order.CreatedAt,
		&order.UpdatedAt,
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

func scanOrderline(rows pgx.Rows, orderline *model.Orderline) error {
	return rows.Scan(
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

func (repo *OrderRepo) GetOrder(ctx context.Context, orderID uuid.UUID) (*model.Order, error) {
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

	query := getFullOrderQuery(orderID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := tx.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullOrder: %w", err)
	}
	defer rows.Close()

	orderMap := make(map[string]*model.Order)
	for rows.Next() {
		order := &model.Order{}
		orderline := &model.Orderline{}

		if err = scanFullOrder(rows, order, orderline); err != nil {
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

	return orderMap[orderID.String()], nil
}

func (repo *OrderRepo) GetOrders(ctx context.Context, searchParams dto.SearchOrderDTO) ([]*model.Order, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Acquire in GetOrders: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetOrders transaction: %w", err)
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

	query := searchOrdersQuery(searchParams)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := tx.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query searchOrders: %w", err)
	}

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

	orders := make([]*model.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}

func createOrderlinesInTx(ctx context.Context, tx pgx.Tx, orderlines []*model.Orderline) error {
	batch := &pgx.Batch{}

	for _, orderline := range orderlines {
		query := createOrderlineQuery(orderline)

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("failed to get sql query: %w", err)
		}

		batch.Queue(sqlQuery, args...)
	}

	batchResults := tx.SendBatch(ctx, batch)
	return batchResults.Close()
}

func (repo *OrderRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in GetOrders: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin GetOrders transaction: %w", err)
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

	query := createOrderQuery(order)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec createOrder: %w", err)
	}

	if err = createOrderlinesInTx(ctx, tx, order.Orderlines); err != nil {
		return fmt.Errorf("failed to create orderlines in tx: %w", err)
	}

	return nil
}

func updateOrderInTx(ctx context.Context, tx pgx.Tx, orderID uuid.UUID) error {
	query := updateOrderQuery(orderID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query")
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateOrder: %w", err)
	}

	return nil
}

func (repo *OrderRepo) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	query := deleteOrderQuery(orderID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err := repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteOrder: %w", err)
	}

	return nil
}

func (repo *OrderRepo) DeleteUserOrders(ctx context.Context, userID uuid.UUID) error {
	query := deleteUserOrdersQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err := repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteUserOrders: %w", err)
	}

	return nil
}

func (repo *OrderRepo) GetOrderline(ctx context.Context, orderID, productID uuid.UUID) (*model.Orderline, error) {
	query := getOrderlineQuery(orderID, productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getOrderline: %w", err)
	}

	orderlineMap := make(map[string]*model.Orderline)
	for rows.Next() {
		orderline := &model.Orderline{}

		if err = scanOrderline(rows, orderline); err != nil {
			return nil, fmt.Errorf("failed to scan orderline: %w", err)
		}

		orderlineMap[orderline.OrderID.String()+orderline.ProductID.String()] = orderline
	}

	return orderlineMap[orderID.String()+productID.String()], nil
}

func (repo *OrderRepo) CreateOrderline(ctx context.Context, orderline *model.Orderline) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in UpdateOrderline: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin UpdateOrderline transaction: %w", err)
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

	query := createOrderlineQuery(orderline)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec createOrderline: %w", err)
	}

	if err = updateOrderInTx(ctx, tx, orderline.OrderID); err != nil {
		return fmt.Errorf("failed to update order in transaction: %w", err)
	}

	return nil
}

func (repo *OrderRepo) UpdateOrderline(ctx context.Context, orderline *model.Orderline) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in UpdateOrderline: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin UpdateOrderline transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if err = tx.Rollback(ctx); err != nil {
				repo.logger.Error("failed to rollback transaction", err)
			}
		} else {
			if err = tx.Commit(ctx); err != nil {
				repo.logger.Error("failed to commit transaction", err)
			}
		}
	}()

	query := updateOrderlineQuery(orderline)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateOrderline: %w", err)
	}

	if err = updateOrderInTx(ctx, tx, orderline.OrderID); err != nil {
		return fmt.Errorf("failed to update Order in Tx: %w", err)
	}

	return nil
}

func (repo *OrderRepo) DeleteOrderline(ctx context.Context, orderID, productID uuid.UUID) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in DeleteOrderline: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin DeleteOrderline transaction: %w", err)
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

	query := deleteOrderlineQuery(orderID, productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteOrderline: %w", err)
	}

	if err = updateOrderInTx(ctx, tx, orderID); err != nil {
		return fmt.Errorf("failed to update order in transaction: %w", err)
	}

	return nil
}
