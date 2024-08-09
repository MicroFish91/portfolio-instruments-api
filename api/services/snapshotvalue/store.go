package snapshotvalue

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSnapshotValueStore struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresSnapshotValueStore(db *pgxpool.Pool, logger *slog.Logger) *PostgresSnapshotValueStore {
	return &PostgresSnapshotValueStore{
		db:     db,
		logger: logger,
	}
}
