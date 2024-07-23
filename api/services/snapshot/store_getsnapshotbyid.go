package snapshot

import (
	"context"
	"database/sql"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GetSnapshotById(ctx context.Context, snapshotId, userId int) (*types.Snapshot, *[]types.SnapshotValues, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM snapshots
		WHERE user_id = $1
		AND snap_id = $2`,
		userId, snapshotId,
	)

	snapshot, err := s.parseRowIntoSnapshot(row)
	if err != nil {
		return nil, nil, err
	}

	rows, err := s.db.Query(
		c,
		`SELECT * FROM snapshots_values
		WHERE user_id = $1
		AND snap_id = $2
		ORDER BY account_id ASC, holding_id ASC`,
		userId, snapshotId,
	)

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	snapshotValues, err := s.parseRowsIntoSnapshotValues(rows)

	if err != nil {
		return nil, nil, err
	}
	return snapshot, snapshotValues, nil
}

func (s *PostgresSnapshotStore) parseRowIntoSnapshot(row pgx.Row) (*types.Snapshot, error) {
	var snap types.Snapshot
	var benchmark_id sql.NullInt64

	err := row.Scan(
		&snap.Snap_id,
		&snap.Description,
		&snap.Snap_date,
		&snap.Total,
		&benchmark_id,
		&snap.User_id,
		&snap.Created_at,
		&snap.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	if benchmark_id.Valid {
		snap.Benchmark_id = int(benchmark_id.Int64)
	} else {
		snap.Benchmark_id = 0
	}

	return &snap, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoSnapshotValues(rows pgx.Rows) (*[]types.SnapshotValues, error) {
	var snapshotValues []types.SnapshotValues
	for rows.Next() {
		var sv types.SnapshotValues
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

	return &snapshotValues, nil
}
