package holding

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) DeleteHolding(ctx context.Context, userId, holdingId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				holdings
			where
				holding_id = $1
				and user_id = $2
			returning
				*
		`,
		holdingId, userId,
	)

	holding, err := s.parseRowIntoHolding(row)
	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
