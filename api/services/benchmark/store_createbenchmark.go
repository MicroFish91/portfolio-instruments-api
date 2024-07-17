package benchmark

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) CreateBenchmark(ctx context.Context, b *types.Benchmark) error {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := s.db.Exec(
		c,
		`INSERT INTO benchmarks
		(name, description, asset_allocation, std_dev_pct, real_return_pct, drawdown_yrs, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		b.Name, b.Description, b.Asset_allocation, b.Std_dev_pct, b.Real_return_pct, b.Drawdown_yrs, b.Is_deprecated, b.User_id,
	)

	if err != nil {
		return err
	}

	return nil
}
