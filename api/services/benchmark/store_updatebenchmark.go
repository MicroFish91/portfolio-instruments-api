package benchmark

import (
	"context"
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) UpdateBenchmark(ctx context.Context, b *types.Benchmark) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if b == nil {
		return types.Benchmark{}, errors.New("service error: benchmark struct cannot be nil, valid benchmark data is required")
	}

	row := s.db.QueryRow(
		c,
		`
			update
				benchmarks
			set
				name = $1,
				description = $2,
				asset_allocation = $3,
				std_dev_pct = $4,
				real_return_pct = $5,
				drawdown_yrs = $6,
				is_deprecated = $7,
				rebalance_threshold_pct = $10,
				updated_at = now()
			where
				benchmark_id = $8
				and user_id = $9
			returning
				*
		`,
		b.Name,
		b.Description,
		b.Asset_allocation,
		b.Std_dev_pct,
		b.Real_return_pct,
		b.Drawdown_yrs,
		b.Is_deprecated,
		b.Benchmark_id,
		b.User_id,
		b.Rebalance_threshold_pct,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)
	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, nil
}
