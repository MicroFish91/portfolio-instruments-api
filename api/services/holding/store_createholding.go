package holding

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) CreateHolding(ctx context.Context, h types.Holding) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO holdings
		(name, ticker, asset_category, expense_ratio_pct, maturation_date, interest_rate_pct, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *`,
		h.Name, h.Ticker, h.Asset_category, h.Expense_ratio_pct, h.Maturation_date, h.Interest_rate_pct, h.Is_deprecated, h.User_id,
	)

	holding, err := s.parseRowIntoHolding(row)
	if err != nil {
		return types.Holding{}, err
	}

	return holding, nil
}
