package account

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAccountStore struct {
	db     *pgxpool.Pool // Todo swap with pgxpool
	logger *slog.Logger
}

func NewPostgresAccountStore(postgresDb *pgxpool.Pool, logger *slog.Logger) *PostgresAccountStore {
	return &PostgresAccountStore{
		db:     postgresDb,
		logger: logger,
	}
}
