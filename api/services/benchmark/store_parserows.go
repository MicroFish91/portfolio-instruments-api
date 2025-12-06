package benchmark

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

const benchmarkColumns = `
	benchmark_id,
    name,
    description,
    asset_allocation,
    std_dev_pct,
    real_return_pct,
    drawdown_yrs,
    rec_rebalance_threshold_pct,
    is_deprecated,
    user_id,
    created_at,
    updated_at
`

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
		&b.Rec_rebalance_threshold_pct,
		&b.Is_deprecated,
		&b.User_id,
		&b.Created_at,
		&b.Updated_at,
	)

	if err != nil {
		return types.Benchmark{}, err
	}
	return b, nil
}

func (s *PostgresBenchmarkStore) parseRowsIntoBenchmarks(rows pgx.Rows) ([]types.Benchmark, int, error) {
	var total_items int
	var benchmarks []types.Benchmark

	for rows.Next() {
		var b types.Benchmark
		err := rows.Scan(
			&b.Benchmark_id,
			&b.Name,
			&b.Description,
			&b.Asset_allocation,
			&b.Std_dev_pct,
			&b.Real_return_pct,
			&b.Drawdown_yrs,
			&b.Rec_rebalance_threshold_pct,
			&b.Is_deprecated,
			&b.User_id,
			&b.Created_at,
			&b.Updated_at,
			&total_items,
		)

		if err != nil {
			return nil, 0, err
		}
		benchmarks = append(benchmarks, b)
	}

	return benchmarks, total_items, nil
}
