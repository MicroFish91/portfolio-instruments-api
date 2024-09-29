package benchmark

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresBenchmarkStore) GetBenchmarkByName(ctx context.Context, name string, userId int) (types.Benchmark, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	namePattern := fmt.Sprintf("^%s$", name)

	row := s.db.QueryRow(
		c,
		`SELECT * FROM benchmarks
		WHERE user_id = $1
		AND name ~* $2
		AND is_deprecated = false`,
		userId, namePattern,
	)

	benchmark, err := s.parseRowIntoBenchmark(row)

	if err != nil {
		return types.Benchmark{}, err
	}
	return benchmark, err
}
