package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) RefreshSnapshotTotal(ctx context.Context, userId, snapshotId int) (float64, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_LONG)
	defer cancel()

	// Use an aggregate function to sum row totals
	snapshotTotal, err := s.GetSnapshotTotal(c, userId, snapshotId, &types.GetSnapshotTotalStoreOptions{
		Omit_skip_reb: false,
	})

	if err != nil {
		return 0, err
	}

	// Update snapshots with the new total
	_, err = s.db.Exec(
		c,
		`
			update
				snapshots
			set 
				total = $1
			where
				user_id = $2
				and snap_id = $3`,
		snapshotTotal,
		userId,
		snapshotId,
	)

	return snapshotTotal, err
}
