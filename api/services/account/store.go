package account

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PostgresAccountStore struct {
	db     *pgx.Conn // Todo swap with pgxpool
	logger *slog.Logger
}

func NewPostgresAccountStore(postgresDb *pgx.Conn, logger *slog.Logger) *PostgresAccountStore {
	return &PostgresAccountStore{
		db:     postgresDb,
		logger: logger,
	}
}
