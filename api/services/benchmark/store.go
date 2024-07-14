package benchmark

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PostgresBenchmarkStore struct {
	db     *pgx.Conn
	logger *slog.Logger
}

func NewPostgresBenchmarkStore(db *pgx.Conn, logger *slog.Logger) *PostgresBenchmarkStore {
	return &PostgresBenchmarkStore{
		db:     db,
		logger: logger,
	}
}
