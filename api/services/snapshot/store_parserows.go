package snapshot

import (
	"database/sql"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/jackc/pgx/v5"
)

const snapshotColumns = `
	snap_id,
	description,
	snap_date,
	total,
	weighted_er_pct,
	rebalance_threshold_pct,
	value_order,
	benchmark_id,
	user_id,
	created_at,
	updated_at
`

func (s *PostgresSnapshotStore) parseRowIntoSnapshot(row pgx.Row) (types.Snapshot, error) {
	var snap types.Snapshot
	var benchmark_id, rebalance_threshold_pct sql.NullInt64

	err := row.Scan(
		&snap.Snap_id,
		&snap.Description,
		&snap.Snap_date,
		&snap.Total,
		&snap.Weighted_er_pct,
		&rebalance_threshold_pct,
		&snap.Value_order,
		&benchmark_id,
		&snap.User_id,
		&snap.Created_at,
		&snap.Updated_at,
	)

	if err != nil {
		return types.Snapshot{}, err
	}

	snap.Benchmark_id = utils.ConvertNullIntToInt(benchmark_id)
	snap.Rebalance_threshold_pct = utils.ConvertNullIntToInt(rebalance_threshold_pct)

	return snap, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoSnapshots(rows pgx.Rows) ([]types.Snapshot, int, error) {
	var snapshots []types.Snapshot
	var total_items int

	for rows.Next() {
		var s types.Snapshot
		var benchmark_id sql.NullInt64
		var rebalance_threshold_pct sql.NullInt64

		err := rows.Scan(
			&s.Snap_id,
			&s.Description,
			&s.Snap_date,
			&s.Total,
			&s.Weighted_er_pct,
			&rebalance_threshold_pct,
			&s.Value_order,
			&benchmark_id,
			&s.User_id,
			&s.Created_at,
			&s.Updated_at,
			&total_items,
		)

		if err != nil {
			return nil, 0, err
		}

		if rebalance_threshold_pct.Valid {
			s.Rebalance_threshold_pct = int(rebalance_threshold_pct.Int64)
		} else {
			s.Rebalance_threshold_pct = 0
		}

		if benchmark_id.Valid {
			s.Benchmark_id = int(benchmark_id.Int64)
		} else {
			s.Benchmark_id = 0
		}

		snapshots = append(snapshots, s)
	}

	return snapshots, total_items, nil
}
