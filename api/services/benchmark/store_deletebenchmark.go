package benchmark

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) DeleteBenchmark(ctx context.Context, userId, benchmarkId int) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				benchmarks
			where
				benchmark_id = $1
				and user_id = $2
			returning
				*
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
