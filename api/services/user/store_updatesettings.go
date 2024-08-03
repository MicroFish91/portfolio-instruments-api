package user

import (
	"context"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) UpdateSettings(ctx context.Context, settings types.Settings) (types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("UPDATE settings")

	setColumnsQuery := "SET reb_thresh_pct = $x, updated_at = NOW()"
	setColumnsArgs := []any{settings.Reb_thresh_pct}
	if settings.Benchmark_id != 0 {
		setColumnsQuery = fmt.Sprintf("%s, benchmark_id = $x", setColumnsQuery)
		setColumnsArgs = append(setColumnsArgs, settings.Benchmark_id)
	}

	pgxb.AddQueryWithPositionals(setColumnsQuery, setColumnsArgs)
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{settings.User_id})
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
