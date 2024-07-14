package user

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PostgresUserStore struct {
	db     *pgx.Conn // Todo swap with pgxpool
	logger *slog.Logger
}

func NewPostgresUserStore(postgresDb *pgx.Conn, logger *slog.Logger) *PostgresUserStore {
	return &PostgresUserStore{
		db:     postgresDb,
		logger: logger,
	}
}
