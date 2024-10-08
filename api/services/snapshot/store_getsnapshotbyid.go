package snapshot

import (
	"context"
	"database/sql"
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GetSnapshotById(ctx context.Context, snapshotId, userId int) (types.Snapshot, []types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_LONG)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select 
				* 
			from 
				snapshots
			where 
				user_id = $1
				and snap_id = $2
		`,
		userId,
		snapshotId,
	)

	snapshot, err := s.parseRowIntoSnapshot(row)
	if err != nil {
		return types.Snapshot{}, nil, err
	}
	if snapshot.Snap_id == 0 {
		return types.Snapshot{}, nil, errors.New("snapshot not found")
	}

	rows, err := s.db.Query(
		c,
		`
			select 
				* 
			from 
				snapshots_values
			where 
				user_id = $1
				and snap_id = $2
			order by 
				account_id ASC, 
				holding_id ASC
		`,
		userId,
		snapshotId,
	)

	if err != nil {
		return types.Snapshot{}, nil, err
	}
	defer rows.Close()

	snapshotValues, err := s.parseRowsIntoSnapshotValues(rows)

	if err != nil {
		return types.Snapshot{}, nil, err
	}
	return snapshot, snapshotValues, nil
}

func (s *PostgresSnapshotStore) parseRowIntoSnapshot(row pgx.Row) (types.Snapshot, error) {
	var snap types.Snapshot
	var benchmark_id sql.NullInt64

	err := row.Scan(
		&snap.Snap_id,
		&snap.Description,
		&snap.Snap_date,
		&snap.Total,
		&snap.Weighted_er_pct,
		&benchmark_id,
		&snap.User_id,
		&snap.Created_at,
		&snap.Updated_at,
	)

	if err != nil {
		return types.Snapshot{}, err
	}

	if benchmark_id.Valid {
		snap.Benchmark_id = int(benchmark_id.Int64)
	} else {
		snap.Benchmark_id = 0
	}

	return snap, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoSnapshotValues(rows pgx.Rows) ([]types.SnapshotValue, error) {
	var snapshotValues []types.SnapshotValue
	for rows.Next() {
		var sv types.SnapshotValue
		err := rows.Scan(
			&sv.Snap_val_id,
			&sv.Snap_id,
			&sv.Account_id,
			&sv.Holding_id,
			&sv.Total,
			&sv.Skip_rebalance,
			&sv.User_id,
			&sv.Created_at,
			&sv.Updated_at,
		)

		if err != nil {
			return nil, err
		}
		snapshotValues = append(snapshotValues, sv)
	}

	return snapshotValues, nil
}
