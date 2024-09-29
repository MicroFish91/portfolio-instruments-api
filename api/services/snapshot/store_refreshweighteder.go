package snapshot

import (
	"context"
	"math"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) RefreshSnapshotWeightedER(ctx context.Context, userId, snapId int) (weightedER float64, e error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_LONG)
	defer cancel()

	// Pre-sum the holding values into a separate table first, then cross join so that we have access to the total when aggregating the final value.
	// Trying to do it all in one line seems impossible due to not being able to use aggregate function calls containing window function calls
	// e.g. (sv.total / SUM(sv.total) OVER ()) * holdings.expense_ratio AS weighted_expense_ratio

	row := s.db.QueryRow(
		c,
		`
			select 
				sum((sv.total / snapshot_total.total) * holdings.expense_ratio_pct) as weighted_expense_ratio
			from 
				snapshots_values sv
			inner join 
				holdings on holdings.holding_id = sv.holding_id 
			cross join
				(
					select 
						sum(sv_inner.total) as total
					from 
						snapshots_values sv_inner
					where 
						sv_inner.user_id = $1 
						and sv_inner.snap_id = $2
				) as snapshot_total
			where 
				sv.user_id = $3
				and sv.snap_id = $4
		`,
		userId, snapId, userId, snapId,
	)

	expenseRatio, err := s.parseRowIntoWeightedER(row)
	if err != nil {
		return 0, err
	}

	roundedER := math.Round(expenseRatio*1000) / 1000
	_, err = s.db.Exec(
		c,
		`UPDATE snapshots
		SET weighted_er_pct = $1
		WHERE user_id = $2
		AND snap_id = $3`,
		roundedER, userId, snapId,
	)

	if err != nil {
		return 0, err
	}
	return roundedER, nil
}

func (s *PostgresSnapshotStore) parseRowIntoWeightedER(row pgx.Row) (weightedER float64, e error) {
	var expenseRatio float64
	err := row.Scan(
		&expenseRatio,
	)

	if err != nil {
		return 0, err
	}
	return expenseRatio, nil
}
