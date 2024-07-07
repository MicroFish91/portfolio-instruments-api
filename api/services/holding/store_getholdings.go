package holding

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldings(userId int) (*[]types.Holding, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT *
		FROM holdings
		WHERE user_id = $1`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	var holdings []types.Holding
	for rows.Next() {
		var h types.Holding
		if err = rows.Scan(&h.Holding_id, &h.Name, &h.Ticker, &h.Asset_category, &h.Expense_ratio, &h.Is_deprecated, &h.User_id, &h.Created_at, &h.Updated_at); err != nil {
			return nil, err
		}
		holdings = append(holdings, h)
	}
	return &holdings, nil
}
