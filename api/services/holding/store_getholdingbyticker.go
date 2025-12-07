package holding

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldingByTicker(ctx context.Context, ticker string, userId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select
				%s
			from
				holdings
			where
				user_id = $1
				and ticker = $2
		`, holdingsColumns),
		userId, ticker,
	)

	holding, err := s.parseRowIntoHolding(row)

	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
