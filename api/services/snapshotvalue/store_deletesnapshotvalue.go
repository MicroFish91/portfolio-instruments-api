package snapshotvalue

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) DeleteSnapshotValue(ctx context.Context, snapValId, userId int) (types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				snapshots_values
			where
				snap_val_id = $1
				and user_id = $2
			returning
				*
		`,
		snapValId,
		userId,
	)

	snapshotValue, err := s.parseRowIntoSnapshotValue(row)
	if err != nil {
		return types.SnapshotValue{}, err
	}
	return snapshotValue, nil
}
