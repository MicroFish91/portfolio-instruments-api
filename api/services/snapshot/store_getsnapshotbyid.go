package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) GetSnapshotById(ctx context.Context, snapshotId, userId int) (types.Snapshot, []types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_LONG)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select 
				%s 
			from 
				snapshots
			where 
				user_id = $1
				and snap_id = $2
		`, snapshotColumns),
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
		fmt.Sprintf(`
			select 
				%s 
			from 
				snapshots_values
			where 
				user_id = $1
				and snap_id = $2
			order by 
				account_id ASC, 
				holding_id ASC
		`, snapshotValueColumns),
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
