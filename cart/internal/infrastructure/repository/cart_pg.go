package repository

import (
	"context"
	"fmt"
	"time"

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
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
}

func scanFullCart(rows pgx.Rows, cart *model.Cart, cartline *model.CartLine) error {
	return rows.Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
		&cartline.ID,
		&cartline.CartID,
		&cartline.ProductID,
		&cartline.Name,
		&cartline.Quantity,
		&cartline.CreatedAt,
		&cartline.UpdatedAt,
	)
}

func getUserCartInTx(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (*model.Cart, error) {
	rows, err := tx.Query(ctx, getUserCart, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUserCart: %w", err)
	}

	cart := &model.Cart{}
	found := false
	for rows.Next() {
		err := scanCart(rows, cart)
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

func getCartInTx(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*model.Cart, error) {
	rows, err := tx.Query(ctx, getCartByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getCartByID: %w", err)
	}

	cart := &model.Cart{}
	found := false
	for rows.Next() {
		err := scanCart(rows, cart)
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

func (repo *CartRepo) GetCart(ctx context.Context, id uuid.UUID) (*model.Cart, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Acquire in GetCart: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetCart transaction: %w", err)
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

	rows, err := tx.Query(ctx, getFullCartByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullCartByID: %w", err)
	}
	defer rows.Close()

	cartMap := make(map[string]*model.Cart)
	for rows.Next() {
		cart := &model.Cart{}
		cartline := &model.CartLine{}

		err := scanFullCart(rows, cart, cartline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan full cart: %w", err)
		}

		if oldCart, ok := cartMap[cart.ID.String()]; ok {
			cart = oldCart
		} else {
			cart.Cartlines = make([]*model.CartLine, 0)
		}

		cart.Cartlines = append(cart.Cartlines, cartline)
		cartMap[cart.ID.String()] = cart
	}

	cart := cartMap[id.String()]

	if cart == nil {
		cart, err = getCartInTx(ctx, tx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get cart in transaction: %w", err)
		}
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

	rows, err := tx.Query(ctx, getFullUserCart, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getFullUserCart: %w", err)
	}
	defer rows.Close()

	cartMap := make(map[string]*model.Cart)
	for rows.Next() {
		cart := &model.Cart{}
		cartline := &model.CartLine{}

		err := scanFullCart(rows, cart, cartline)
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
	_, err := repo.pg.Pool.Exec(
		ctx,
		createCart,
		cart.ID,
		cart.UserID,
		cart.CreatedAt,
		cart.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createCart: %w", err)
	}

	return nil
}

func (repo *CartRepo) CreateCartline(ctx context.Context, cartline model.CartLine) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		createCartline,
		cartline.ID,
		cartline.CartID,
		cartline.ProductID,
		cartline.Name,
		cartline.Quantity,
		cartline.CreatedAt,
		cartline.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createCartline: %w", err)
	}

	return nil
}

func (repo *CartRepo) UpdateCartline(ctx context.Context, cartline model.CartLine) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		updateCartline,
		cartline.Name,
		cartline.Quantity,
		cartline.UpdatedAt,
		cartline.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec updateCartline: %w", err)
	}

	_, err = repo.pg.Pool.Exec(
		ctx,
		updateCart,
		time.Now(),
		cartline.CartID,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec updateCart: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCart(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteCart, id)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteCart: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCartline(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteCartline, id)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteCartline: %w", err)
	}

	return nil
}

func (repo *CartRepo) DeleteCartCartlines(ctx context.Context, cartID uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteCartCartlines, cartID)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteCartCartlines: %w", err)
	}

	return nil
}
