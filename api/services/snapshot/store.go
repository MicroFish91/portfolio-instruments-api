package snapshot

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PostgresSnapshotStore struct {
	db     *pgx.Conn
	logger *slog.Logger
}

func NewPostgresSnapshotStore(db *pgx.Conn, logger *slog.Logger) *PostgresSnapshotStore {
	return &PostgresSnapshotStore{
		db:     db,
		logger: logger,
	}
}
