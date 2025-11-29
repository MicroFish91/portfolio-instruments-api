package benchmark

import (
	"context"
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) CreateBenchmark(ctx context.Context, b *types.Benchmark) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if b == nil {
		return types.Benchmark{}, errors.New("service error: benchmark struct cannot be nil, valid benchmark data is required")
	}

	var rec_rebalance_threshold_pct int = b.Rec_rebalance_threshold_pct
	if b.Rec_rebalance_threshold_pct == 0 {
		rec_rebalance_threshold_pct = constants.BENCHMARK_REBALANCE_PCT_DEFAULT
	}

	row := s.db.QueryRow(
		c,
		`INSERT INTO benchmarks
		(name, description, asset_allocation, std_dev_pct, real_return_pct, drawdown_yrs, is_deprecated, user_id, rec_rebalance_threshold_pct)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING *`,
		b.Name, b.Description, b.Asset_allocation, b.Std_dev_pct, b.Real_return_pct, b.Drawdown_yrs, b.Is_deprecated, b.User_id, rec_rebalance_threshold_pct,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)

	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, nil
}
