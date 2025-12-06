package benchmark

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) GetBenchmarkById(ctx context.Context, userId, benchmarkId int) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select
				%s 
			from 
				benchmarks
			where 
				benchmark_id = $1
				and user_id = $2
		`, benchmarkColumns),
		benchmarkId, userId,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)

	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, nil
}
