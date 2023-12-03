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

func (repo *UserRepo) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	query := getUserQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUser: %w", err)
	}

	userMap := make(map[string]*model.User)
	for rows.Next() {
		user := &model.User{}

		if err = scanUser(rows, user); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		userMap[user.ID.String()] = user
	}

	return userMap[userID.String()], nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := getUserByEmailQuery(email)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUserByEmail: %w", err)
	}

	userMap := make(map[string]*model.User)
	for rows.Next() {
		user := &model.User{}

		if err = scanUser(rows, user); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		userMap[user.Email] = user
	}

	return userMap[email], nil
}

func (repo *UserRepo) GetUsers(ctx context.Context) ([]*model.User, error) {
	query := getUsersQuery()

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql query: %w", err)
	}

	rows, err := repo.pg.Pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to Query getUsers: %w", err)
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}
		if err = scanUser(rows, user); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	query := createUserQuery(user)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec createUser: %w", err)
	}

	return nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, user *model.User) error {
	query := updateUserQuery(user)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec updateUser: %w", err)
	}

	return nil
}

func (repo *UserRepo) ChangeUserRole(ctx context.Context, userID uuid.UUID, role model.UserRoles) error {
	query := changeUserRoleQuery(userID, role)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec changeUserRole: %w", err)
	}

	return nil
}

func (repo *UserRepo) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := deleteUserQuery(userID)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to get sql query: %w", err)
	}

	if _, err = repo.pg.Pool.Exec(ctx, sqlQuery, args...); err != nil {
		return fmt.Errorf("failed to Exec deleteUser: %w", err)
	}

	return nil
}
