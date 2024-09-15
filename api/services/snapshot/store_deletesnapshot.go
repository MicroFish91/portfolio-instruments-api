package snapshot

import (
	"context"
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) DeleteSnapshot(ctx context.Context, snapId, userId int) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				snapshots
			where
				snap_id = $1
				and user_id = $2
			returning
				*
		`,
		snapId,
		userId,
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
