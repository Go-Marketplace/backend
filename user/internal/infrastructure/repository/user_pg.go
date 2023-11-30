package repository

import (
	"context"
	"fmt"

	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
	"github.com/Go-Marketplace/backend/user/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	pg     *postgres.Postgres
	logger *logger.Logger
}

func NewUserRepo(pg *postgres.Postgres, logger *logger.Logger) *UserRepo {
	return &UserRepo{
		pg:     pg,
		logger: logger,
	}
}

func scanUser(rows pgx.Rows, user *model.User) error {
	return rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.Address,
		&user.Phone,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (repo *UserRepo) GetUser(ctx context.Context, id uuid.UUID) (*model.User, error) {
	rows, err := repo.pg.Pool.Query(ctx, getUserByID, id)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUserByID: %w", err)
	}

	user := &model.User{}
	found := false
	for rows.Next() {
		err := scanUser(rows, user)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		found = true
	}

	if !found {
		return nil, nil
	}

	return user, nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	rows, err := repo.pg.Pool.Query(ctx, getUserByEmail, email)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUserByEmail: %w", err)
	}

	user := &model.User{}
	found := false
	for rows.Next() {
		err := scanUser(rows, user)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		found = true
	}

	if !found {
		return nil, nil
	}

	return user, nil
}

func (repo *UserRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	conn, err := repo.pg.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Acquire in GetAllUsers: %w", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to begin GetAllUsers transaction: %w", err)
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

	rows, err := tx.Query(ctx, getAllUsers)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getAllUsers: %w", err)
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}
		err := scanUser(rows, user)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepo) CreateUser(ctx context.Context, user model.User) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		createUser,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
		user.Address,
		user.Phone,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec createUser: %w", err)
	}

	return nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, user model.User) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		updateUser,
		user.FirstName,
		user.LastName,
		user.Address,
		user.Phone,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec updateUser: %w", err)
	}

	return nil
}

func (repo *UserRepo) ChangeUserRole(ctx context.Context, user model.User) error {
	_, err := repo.pg.Pool.Exec(
		ctx,
		changeUserRole,
		user.Role,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to Exec changeUserRole: %w", err)
	}

	return nil
}

func (repo *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := repo.pg.Pool.Exec(ctx, deleteUser, id)
	if err != nil {
		return fmt.Errorf("failed to Exec deleteUser: %w", err)
	}

	return nil
}
