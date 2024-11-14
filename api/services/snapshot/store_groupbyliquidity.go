package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
)

func (s *PostgresSnapshotStore) GroupByLiquidity(ctx context.Context, userId, snapId int) (float64, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				SUM(snapshots_values.total) AS liquid_total
			from
				snapshots_values
			inner join
				accounts
			on
				snapshots_values.account_id = accounts.account_id
			inner join
				holdings
			on
				snapshots_values.holding_id = holdings.holding_id
			where
				snapshots_values.user_id = $1
				and snapshots_values.snap_id = $2
				and accounts.tax_shelter = 'TAXABLE'
				and holdings.asset_category = 'CASH'
		`,
		userId, snapId,
	)

	var liquid_total float64
	err := row.Scan(
		&liquid_total,
	)

	if err != nil {
		return 0, err
	}
	return liquid_total, nil
}
