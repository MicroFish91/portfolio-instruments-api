package holding

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresHoldingStore) GetHoldingById(ctx context.Context, userId, holdingId int) (types.Holding, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT *
		FROM holdings
		WHERE user_id = $1 AND holding_id = $2`,
		userId, holdingId,
	)

	holding, err := s.parseRowIntoHolding(row)
	if err != nil {
		return types.Holding{}, err
	}

	return holding, nil
}

func (s *PostgresHoldingStore) parseRowIntoHolding(row pgx.Row) (types.Holding, error) {
	var holding types.Holding
	err := row.Scan(
		&holding.Holding_id,
		&holding.Name,
		&holding.Ticker,
		&holding.Asset_category,
		&holding.Expense_ratio_pct,
		&holding.Maturation_date,
		&holding.Interest_rate_pct,
		&holding.Is_deprecated,
		&holding.User_id,
		&holding.Created_at,
		&holding.Updated_at,
	)

	if err != nil {
		return types.Holding{}, err
	}
	return holding, nil
}
