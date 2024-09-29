package snapshotvalue

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) DeleteSnapshotValue(ctx context.Context, snapId, snapValId, userId int) (types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				snapshots_values
			where
				snap_id = $1 
				and snap_val_id = $2
				and user_id = $3
			returning
				*
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
