package holding

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) CreateHolding(ctx context.Context, h types.Holding) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO holdings
		(name, ticker, asset_category, expense_ratio, maturation_date, interest_rate, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *`,
		h.Name, h.Ticker, h.Asset_category, h.Expense_ratio, h.Maturation_date, h.Interest_rate, h.Is_deprecated, h.User_id,
	)

	holding, err := s.parseRowIntoHolding(row)
	if err != nil {
		return types.Holding{}, err
	}

	return holding, nil
}
