package snapshot

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) TallyByWeightedER(ctx context.Context, userId, snapId int) (weightedER float64, e error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select 
				sum((sv.total / snapshot_total.total) * holdings.expense_ratio) as weighted_expense_ratio
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

	return expenseRatio, nil
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
