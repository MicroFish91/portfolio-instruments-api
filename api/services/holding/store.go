package holding

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresHoldingStore struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresHoldingStore(db *pgxpool.Pool, logger *slog.Logger) *PostgresHoldingStore {
	return &PostgresHoldingStore{
		db:     db,
		logger: logger,
	}
}
