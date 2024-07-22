package settings

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresSettingsStore struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresSettingsStore(db *pgxpool.Pool, logger *slog.Logger) *PostgresSettingsStore {
	return &PostgresSettingsStore{
		db:     db,
		logger: logger,
	}
}
