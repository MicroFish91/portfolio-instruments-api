package holding

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldingByName(ctx context.Context, name string, userId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				*
			from
				holdings
			where
				name = $1
				and user_id = $2
				and is_deprecated = false
		`,
		name, userId,
	)

	holding, err := s.parseRowIntoHolding(row)

	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
