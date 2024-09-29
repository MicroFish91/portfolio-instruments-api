package benchmark

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) CreateBenchmark(ctx context.Context, b types.Benchmark) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO benchmarks
		(name, description, asset_allocation, std_dev_pct, real_return_pct, drawdown_yrs, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *`,
		b.Name, b.Description, b.Asset_allocation, b.Std_dev_pct, b.Real_return_pct, b.Drawdown_yrs, b.Is_deprecated, b.User_id,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)

	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, nil
}
