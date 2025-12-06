package holding

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

const holdingsColumns = `
	holding_id,
	name,
	ticker,
	asset_category,
	expense_ratio_pct,
	maturation_date,
	interest_rate_pct,
	is_deprecated,
	user_id,
	created_at,
	updated_at
`

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

func (s *PostgresHoldingStore) parseRowsIntoHoldings(rows pgx.Rows) ([]types.Holding, int, error) {
	var total_items int
	var holdings []types.Holding

	for rows.Next() {
		var h types.Holding
		err := rows.Scan(
			&h.Holding_id,
			&h.Name,
			&h.Ticker,
			&h.Asset_category,
			&h.Expense_ratio_pct,
			&h.Maturation_date,
			&h.Interest_rate_pct,
			&h.Is_deprecated,
			&h.User_id,
			&h.Created_at,
			&h.Updated_at,
			&total_items,
		)

		if err != nil {
			return nil, 0, err
		}
		holdings = append(holdings, h)
	}

	return holdings, total_items, nil
}
