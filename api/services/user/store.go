package user

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserStore struct {
	db     *pgxpool.Pool // Todo swap with pgxpool
	logger *slog.Logger
}

func NewPostgresUserStore(postgresDb *pgxpool.Pool, logger *slog.Logger) *PostgresUserStore {
	return &PostgresUserStore{
		db:     postgresDb,
		logger: logger,
	}
}
