package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) GetSnapshotById(ctx context.Context, snapshotId, userId int) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
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
		return types.Snapshot{}, err
	}
	if snapshot.Snap_id == 0 {
		return types.Snapshot{}, errors.New("snapshot not found")
	}

	return snapshot, nil
}
