package repository

import (
	"time"

	"github.com/Go-Marketplace/backend/user/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func getUsersQuery() sq.SelectBuilder {
	return psql.Select(
		"user_id",
		"first_name",
		"last_name",
		"password",
		"email",
		"address",
		"phone",
		"role",
		"created_at",
		"updated_at",
	).
		From("users")
}

func getUserQuery(userID uuid.UUID) sq.SelectBuilder {
	return getUsersQuery().
		Where(sq.Eq{
			"user_id": userID,
		})
}

func getUserByEmailQuery(email string) sq.SelectBuilder {
	return getUsersQuery().
		Where(sq.Eq{
			"email": email,
		})
}

func createUserQuery(user *model.User) sq.InsertBuilder {
	return psql.Insert("users").
		Columns(
			"user_id",
			"first_name",
			"last_name",
			"password",
			"email",
			"address",
			"phone",
			"role",
			"created_at",
			"updated_at",
		).
		Values(
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
}

func updateUserQuery(user *model.User) sq.UpdateBuilder {
	query := psql.Update("users")

	if user.FirstName != "" {
		query = query.Set("first_name", user.FirstName)
	}

	if user.LastName != "" {
		query = query.Set("last_name", user.LastName)
	}

	if user.Address != "" {
		query = query.Set("address", user.Address)
	}

	if user.Phone != "" {
		query = query.Set("phone", user.Phone)
	}

	return query.Set("updated_at", time.Now()).
		Where(sq.Eq{
			"user_id": user.ID,
		})
}

func changeUserRoleQuery(userID uuid.UUID, role model.UserRoles) sq.UpdateBuilder {
	return psql.Update("users").
		Set("role", role).
		Set("updated_at", time.Now()).
		Where(sq.Eq{
			"user_id": userID,
		})
}

func deleteUserQuery(userID uuid.UUID) sq.DeleteBuilder {
	return psql.Delete("users").
		Where(sq.Eq{
			"user_id": userID,
		})
}
