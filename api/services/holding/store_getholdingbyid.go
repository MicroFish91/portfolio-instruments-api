package holding

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresHoldingStore) GetHoldingById(ctx context.Context, userId int, holdingId int) (*types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT *
		FROM holdings
		WHERE user_id = $1 AND holding_id = $2`,
		userId, holdingId,
	)

	var holding types.Holding
	err := row.Scan(
		&holding.Holding_id,
		&holding.Name,
		&holding.Ticker,
		&holding.Asset_category,
		&holding.Expense_ratio,
		&holding.Maturation_date,
		&holding.Interest_rate,
		&holding.Is_deprecated,
		&holding.User_id,
		&holding.Created_at,
		&holding.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &holding, nil
}
