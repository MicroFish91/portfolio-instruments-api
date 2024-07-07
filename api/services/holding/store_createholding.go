package holding

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) CreateHolding(h *types.Holding) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO holdings
		(name, ticker, asset_category, expense_ratio, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		h.Name, h.Ticker, h.Asset_category, h.Expense_ratio, h.Is_deprecated, h.User_id,
	)

	if err != nil {
		return err
	}
	return nil
}
