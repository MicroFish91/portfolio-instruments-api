package holding

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldingByName(ctx context.Context, name string, userId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	namePattern := fmt.Sprintf("^%s$", name)

	row := s.db.QueryRow(
		c,
		`
			select
				*
			from
				holdings
			where
				name ~* $1
				and user_id = $2
				and is_deprecated = false
		`,
		namePattern, userId,
	)

	holding, err := s.parseRowIntoHolding(row)

	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
