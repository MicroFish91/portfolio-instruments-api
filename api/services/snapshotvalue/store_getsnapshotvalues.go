package snapshotvalue

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) GetSnapshotValues(ctx context.Context, snapId, userId int) ([]types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	rows, err := s.db.Query(
		c,
		fmt.Sprintf(`
			select
				%s
			from
				snapshots_values
			where
				snap_id = $1
				and user_id = $2
		`, snapshotValueColumns),
		snapId,
		userId,
	)

	if err != nil {
		return []types.SnapshotValue{}, err
	}
	defer rows.Close()

	snapshotValues, err := s.parseRowsIntoSnapshotValues(rows)
	if err != nil {
		return []types.SnapshotValue{}, err
	}
	if len(snapshotValues) == 0 {
		return []types.SnapshotValue{}, errors.New("snapshot_values for the given snapshot id not found")
	}

	return snapshotValues, nil
}
