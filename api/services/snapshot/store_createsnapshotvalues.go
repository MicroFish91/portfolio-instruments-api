package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) CreateSnapshotValues(ctx context.Context, snapVals *types.SnapshotValues) (*types.SnapshotValues, error) {
	row := s.db.QueryRow(
		ctx,
		`INSERT INTO snapshots_values
		(snap_id, account_id, holding_id, total, skip_rebalance, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		snapVals.Snap_id, snapVals.Account_id, snapVals.Holding_id, snapVals.Total, snapVals.Skip_rebalance, snapVals.User_id,
	)

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
		return nil, err
	}

	return &sv, nil
}
