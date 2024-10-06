package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresUserStore) CreateSettings(ctx context.Context, set *types.Settings) (types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if set == nil {
		return types.Settings{}, errors.New("internal error: settings struct cannot be nil, valid settings data is required")
	}

	row := s.db.QueryRow(
		c,
		`INSERT INTO settings
		(reb_thresh_pct, user_id)
		VALUES ($1, $2)
		RETURNING *`,
		set.Reb_thresh_pct, set.User_id,
	)

	settings, err := s.parseRowIntoSettings(row)
	if err != nil {
		return types.Settings{}, err
	}

	return settings, nil
}

func (s *PostgresUserStore) parseRowIntoSettings(row pgx.Row) (types.Settings, error) {
	var setting types.Settings
	var benchmark_id sql.NullInt64

	err := row.Scan(
		&setting.Settings_id,
		&setting.Reb_thresh_pct,
		&setting.User_id,
		&benchmark_id,
		&setting.Created_at,
		&setting.Updated_at,
	)

	if err != nil {
		return types.Settings{}, err
	}

	if benchmark_id.Valid {
		setting.Benchmark_id = int(benchmark_id.Int64)
	} else {
		setting.Benchmark_id = 0
	}

	return setting, nil
}
