package settings

import (
	"context"
	"database/sql"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSettingsStore) CreateSettings(ctx context.Context, settings *types.Settings) (*types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO settings
		(reb_thresh_pct, vp_thresh_pct, vp_enabled, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING *`,
		settings.Reb_thresh_pct, settings.Vp_thresh_pct, settings.Vp_enabled, settings.User_id,
	)

	settings, err := s.parseRowIntoSettings(row)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (s *PostgresSettingsStore) parseRowIntoSettings(row pgx.Row) (*types.Settings, error) {
	var setting types.Settings
	var benchmark_id sql.NullInt64

	err := row.Scan(
		&setting.Settings_id,
		&setting.Reb_thresh_pct,
		&setting.Vp_thresh_pct,
		&setting.Vp_enabled,
		&setting.User_id,
		&benchmark_id,
		&setting.Created_at,
		&setting.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	if benchmark_id.Valid {
		setting.Benchmark_id = int(benchmark_id.Int64)
	} else {
		setting.Benchmark_id = 0
	}

	return &setting, nil
}
