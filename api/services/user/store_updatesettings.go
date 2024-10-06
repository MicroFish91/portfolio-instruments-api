package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) UpdateSettings(ctx context.Context, set *types.Settings) (types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if set == nil {
		return types.Settings{}, errors.New("service error: settings struct cannot be nil, valid settings data is required")
	}

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("UPDATE settings")

	setColumnsQuery := "SET reb_thresh_pct = $x, updated_at = NOW()"
	setColumnsArgs := []any{set.Reb_thresh_pct}
	if set.Benchmark_id != 0 {
		setColumnsQuery = fmt.Sprintf("%s, benchmark_id = $x", setColumnsQuery)
		setColumnsArgs = append(setColumnsArgs, set.Benchmark_id)
	}

	pgxb.AddQueryWithPositionals(setColumnsQuery, setColumnsArgs)
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{set.User_id})
	pgxb.AddQuery("RETURNING *")

	row := s.db.QueryRow(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	settings, err := s.parseRowIntoSettings(row)

	if err != nil {
		return types.Settings{}, err
	}
	return settings, nil
}
