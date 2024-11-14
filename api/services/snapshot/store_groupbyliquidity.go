package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GroupByLiquidity(ctx context.Context, userId, snapId int) ([]types.LiquidityResource, float64, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	rows, err := s.db.Query(
		c,
		`
			select
				a.name,
				h.name,
				h.asset_category,
				h.ticker,
				a.institution,
				a.tax_shelter,
				sv.total,
				SUM(sv.total) AS liquid_total
			from
				snapshots_values sv
			inner join
				accounts a on sv.account_id = a.account_id
			inner join
				holdings h on sv.holding_id = h.holding_id
			where
				sv.user_id = $1
				and sv.snap_id = $2
				and a.tax_shelter = 'TAXABLE'
				and h.asset_category = 'CASH'
			group by
				a.institution, a.tax_shelter, a.name, h.ticker, h.name, h.asset_category, sv.total
		`,
		userId, snapId,
	)

	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	resources, liquidTotal, err := s.parseRowsIntoLiquidityResources(rows)
	if err != nil {
		return nil, 0, err
	}

	return resources, liquidTotal, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoLiquidityResources(rows pgx.Rows) ([]types.LiquidityResource, float64, error) {
	var liquid_total float64
	var resources []types.LiquidityResource

	for rows.Next() {
		var r types.LiquidityResource
		err := rows.Scan(
			&r.Account_name,
			&r.Holding_name,
			&r.Asset_category,
			&r.Ticker,
			&r.Institution,
			&r.TaxShelter,
			&r.Total,
			&liquid_total,
		)

		if err != nil {
			return nil, 0, err
		}
		resources = append(resources, r)
	}

	return resources, liquid_total, nil
}
