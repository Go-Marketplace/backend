package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CartRepo struct {
	pg     *postgres.Postgres
	logger *logger.Logger
}

func NewCartRepo(pg *postgres.Postgres, logger *logger.Logger) *CartRepo {
	return &CartRepo{
		pg:     pg,
		logger: logger,
	}
}

func scanCart(rows pgx.Rows, cart *model.Cart) error {
	return rows.Scan(
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
}

func scanFullCart(rows pgx.Rows, cart *model.Cart, cartline *model.CartLine) error {
	return rows.Scan(
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
		&cartline.UserID,
		&cartline.ProductID,
		&cartline.Name,
		&cartline.Quantity,
		&cartline.CreatedAt,
		&cartline.UpdatedAt,
	)
}

func scanCartline(rows pgx.Rows, cartline *model.CartLine) error {
	return rows.Scan(
		&cartline.UserID,
		&cartline.ProductID,
		&cartline.Name,
		&cartline.Quantity,
		&cartline.CreatedAt,
		&cartline.UpdatedAt,
	)
}

func getUserCartInTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (*model.Cart, error) {
	query := getUserCartQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := tx.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUserCart: %w", err)
	}

	cart := &model.Cart{}
	found := false
	for rows.Next() {
		err = scanCart(rows, cart)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cart: %w", err)
		}
		found = true
	}

	if !found {
		return nil, nil
	}

	return cart, nil
}

func (repo *CartRepo) GetUserCart(ctx context.Context, userID uuid.UUID) (*model.Cart, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Acquire in GetUserCart: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetUserCart transaction: %w", err)
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

	query := getFullUserCartQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := tx.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullUserCart: %w", err)
	}
	defer rows.Close()

	cartMap := make(map[string]*model.Cart)
	for rows.Next() {
		cart := &model.Cart{}
		cartline := &model.CartLine{}

		err = scanFullCart(rows, cart, cartline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full cart: %w", err)
		}

		if oldCart, ok := cartMap[cart.UserID.String()]; ok {
			cart = oldCart
		} else {
			cart.Cartlines = make([]*model.CartLine, 0)
		}

		cart.Cartlines = append(cart.Cartlines, cartline)
		cartMap[cart.UserID.String()] = cart
	}

	cart := cartMap[userID.String()]

	if cart == nil {
		cart, err = getUserCartInTx(ctx, tx, userID)
		if err != nil {
			return nil, fmt.Errorf("failed to get user cart in transaction: %w", err)
		}
	}

	return cart, nil
}

func (repo *CartRepo) CreateCart(ctx context.Context, cart model.Cart) error {
	query := createCartQuery(cart)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec createCart: %w", err)
	}

	return nil
}

func updateCartInTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error {
	query := updateCartQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateCart: %w", err)
	}

	return nil
}

func (repo *CartRepo) GetCartline(ctx context.Context, userID, productID uuid.UUID) (*model.CartLine, error) {
	query := getCartlineQuery(userID, productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getCartline: %w", err)
	}

	cartlineMap := make(map[string]*model.CartLine)
	for rows.Next() {
		cartline := &model.CartLine{}
		
		if err = scanCartline(rows, cartline); err != nil {
			return nil, fmt.Errorf("failed to scan cartline: %w", err)
		}

		cartlineMap[cartline.UserID.String()+cartline.ProductID.String()] = cartline
	}

	return cartlineMap[userID.String()+productID.String()], nil
}

func (repo *CartRepo) CreateCartline(ctx context.Context, cartline *model.CartLine) error {
	query := createCartlineQuery(cartline)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec createCartline: %w", err)
	}

	return nil
}

func (repo *CartRepo) CreateCartlines(ctx context.Context, cartlines []*model.CartLine) error {
	batch := &pgx.Batch{}

	for _, cartline := range cartlines {
		query := createCartlineQuery(cartline)

		sqlQuery, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("failed to get sql query: %w", err)
		}

		batch.Queue(sqlQuery, args...)
	}

	batchResults := repo.pg.Pool.SendBatch(ctx, batch)
	return batchResults.Close()
}

func (repo *CartRepo) UpdateCartline(ctx context.Context, cartline model.CartLine) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in UpdateCartline: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin UpdateCartline transaction: %w", err)
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

	query := updateCartlineQuery(cartline)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateCartline: %w", err)
	}

	if err = updateCartInTx(ctx, tx, cartline.UserID); err != nil {
		return fmt.Errorf("failed to update cart in transaction: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCart(ctx context.Context, userID uuid.UUID) error {
	query := deleteCartQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteCart: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCartline(ctx context.Context, userID uuid.UUID, productID uuid.UUID) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in DeleteCartline: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin DeleteCartline transaction: %w", err)
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

	query := deleteCartlineQuery(userID, productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query")
	}

	if _, err = tx.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteCartline: %w", err)
	}

	if err = updateCartInTx(ctx, tx, userID); err != nil {
		return fmt.Errorf("failed to update cart in transaction: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteProductCartlines(ctx context.Context, productID uuid.UUID) error {
	query := deleteProductCartlinesQuery(productID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query")
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteProductCartlines: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCartCartlines(ctx context.Context, userID uuid.UUID) error {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to Acquire in DeleteCartCartlines: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin DeleteCartCartlines transaction: %w", err)
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

	query := deleteCartCartlinesQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	_, err = tx.Exec(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteCartCartlines: %w", err)
	}

	if err = updateCartInTx(ctx, tx, userID); err != nil {
		return fmt.Errorf("failed to update cart in transaction: %w", err)
	}

	return nil
}
