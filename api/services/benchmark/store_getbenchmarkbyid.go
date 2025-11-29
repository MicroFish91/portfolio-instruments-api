package benchmark

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresBenchmarkStore) GetBenchmarkById(ctx context.Context, userId, benchmarkId int) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				* 
			from 
				benchmarks
			where 
				benchmark_id = $1
				and user_id = $2
		`,
		benchmarkId,
		userId,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)

	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, nil
}

func (s *PostgresBenchmarkStore) parseRowIntoBenchmark(row pgx.Row) (types.Benchmark, error) {
	var b types.Benchmark
	err := row.Scan(
		&b.Benchmark_id,
		&b.Name,
		&b.Description,
		&b.Asset_allocation,
		&b.Std_dev_pct,
		&b.Real_return_pct,
		&b.Drawdown_yrs,
		&b.Is_deprecated,
		&b.User_id,
		&b.Created_at,
		&b.Updated_at,
		&b.Rec_rebalance_threshold_pct,
	)

	if err != nil {
		return types.Benchmark{}, err
	}
	return b, nil
}
