package benchmark

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresBenchmarkStore struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresBenchmarkStore(db *pgxpool.Pool, logger *slog.Logger) *PostgresBenchmarkStore {
	return &PostgresBenchmarkStore{
		db:     db,
		logger: logger,
	}
}
