package snapshotvalue

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) UpdateSnapshotValue(ctx context.Context, sv types.SnapshotValue) (types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			update
				snapshots_values
			set
				account_id = $1,
				holding_id = $2,
				total = $3,
				skip_rebalance = $4,
				updated_at = now()
			where
				snap_val_id = $5
				and snap_id = $6
				and user_id = $7
			returning
				*
		`,
		sv.Account_id,
		sv.Holding_id,
		sv.Total,
		sv.Skip_rebalance,
		sv.Snap_val_id,
		sv.Snap_id,
		sv.User_id,
	)

	snapshotValue, err := s.parseRowIntoSnapshotValue(row)
	if err != nil {
		return types.SnapshotValue{}, err
	}

	return snapshotValue, nil
}
