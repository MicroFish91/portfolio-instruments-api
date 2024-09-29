package holding

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldingByTicker(ctx context.Context, ticker string, userId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM holdings
		WHERE user_id = $1
		AND ticker = $2
		AND is_deprecated = false`,
		userId, ticker,
	)

	holding, err := s.parseRowIntoHolding(row)

	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
