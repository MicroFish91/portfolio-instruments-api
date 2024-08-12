package snapshot

import (
	"context"
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
		return types.Snapshot{}, nil
	}
	return snapshot, nil
}