package snapshotvalue

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) GetSnapshotValue(ctx context.Context, snapId, snapValId, userId int) (types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				*
			from
				snapshots_values
			where
				snap_id = $1
				and snap_val_id = $2
				and user_id = $3
		`,
		snapId,
		snapValId,
		userId,
	)

	snapshotValue, err := s.parseRowIntoSnapshotValue(row)
	if err != nil {
		return types.SnapshotValue{}, err
	}
	return snapshotValue, nil
}
