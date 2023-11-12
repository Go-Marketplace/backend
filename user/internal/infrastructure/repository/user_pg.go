package repository

import (
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/pkg/postgres"
)

type UserRepo struct {
	pg     *postgres.Postgres
	logger logger.Logger
}

func NewUserRepo(pg *postgres.Postgres, logger logger.Logger) *UserRepo {
	return &UserRepo{
		pg: pg,
		logger: logger,
	}
}


