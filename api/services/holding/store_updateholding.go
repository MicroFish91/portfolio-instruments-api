package holding

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) UpdateHolding(ctx context.Context, h types.Holding) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			update 
				holdings
			set
				name = $1,
				ticker = $2,
				asset_category = $3,
				expense_ratio_pct = $4,
				maturation_date = $5,
				interest_rate_pct = $6,
				is_deprecated = $7,
				updated_at = now()
			where
				holding_id = $8
				and user_id = $9
			returning
				*
		`,
		h.Name,
		h.Ticker,
		h.Asset_category,
		h.Expense_ratio_pct,
		h.Maturation_date,
		h.Interest_rate_pct,
		h.Is_deprecated,
		h.Holding_id,
		h.User_id,
	)

	holding, err := s.parseRowIntoHolding(row)
	if err != nil {
		return types.Holding{}, err
	}

	return holding, nil
}
