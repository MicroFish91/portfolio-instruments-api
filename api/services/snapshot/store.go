package snapshot

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSnapshotStore struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresSnapshotStore(db *pgxpool.Pool, logger *slog.Logger) *PostgresSnapshotStore {
	return &PostgresSnapshotStore{
		db:     db,
		logger: logger,
	}
}
