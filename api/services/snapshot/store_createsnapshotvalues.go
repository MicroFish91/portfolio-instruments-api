package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) CreateSnapshotValues(ctx context.Context, snapVals types.SnapshotValues) (types.SnapshotValues, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO snapshots_values
		(snap_id, account_id, holding_id, total, skip_rebalance, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		snapVals.Snap_id, snapVals.Account_id, snapVals.Holding_id, snapVals.Total, snapVals.Skip_rebalance, snapVals.User_id,
	)

	sv, err := s.parseRowIntoSnapshotValues(row)

	if err != nil {
		return types.SnapshotValues{}, err
	}
	return sv, nil
}

func (s *PostgresSnapshotStore) parseRowIntoSnapshotValues(row pgx.Row) (types.SnapshotValues, error) {
	var sv types.SnapshotValues
	err := row.Scan(
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
		return types.SnapshotValues{}, err
	}
	return sv, nil
}
