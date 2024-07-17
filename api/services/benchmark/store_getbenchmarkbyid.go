package benchmark

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) GetBenchmarkById(ctx context.Context, userId, benchmarkId int) (*types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM benchmarks
		WHERE benchmark_id = $1
		AND user_id = $2`,
		benchmarkId, userId,
	)

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
	)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
