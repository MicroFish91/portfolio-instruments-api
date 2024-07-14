package holding

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PostgresHoldingStore struct {
	db     *pgx.Conn
	logger *slog.Logger
}

func NewPostgresHoldingStore(db *pgx.Conn, logger *slog.Logger) *PostgresHoldingStore {
	return &PostgresHoldingStore{
		db:     db,
		logger: logger,
	}
}
