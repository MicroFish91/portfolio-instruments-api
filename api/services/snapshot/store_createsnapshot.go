package snapshot

import (
	"context"
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) CreateSnapshot(ctx context.Context, snapshot *types.Snapshot) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if snapshot == nil {
		return types.Snapshot{}, errors.New("service error: snapshot struct cannot be nil, valid snapshot data is required")
	}

	row := s.db.QueryRow(
		c,
		`INSERT INTO snapshots
		(snap_date, description, user_id)
		VALUES ($1, $2, $3)
		RETURNING *`,
		snapshot.Snap_date, snapshot.Description, snapshot.User_id,
	)

	snap, err := s.parseRowIntoSnapshot(row)
	if err != nil {
		return types.Snapshot{}, err
	}

	if snapshot.Benchmark_id != 0 {
		row = s.db.QueryRow(
			c,
			`UPDATE snapshots
			SET benchmark_id = $1
			WHERE snap_id = $2
			RETURNING *`,
			snapshot.Benchmark_id, snap.Snap_id,
		)

		snap, err = s.parseRowIntoSnapshot(row)
		if err != nil {
			return types.Snapshot{}, err
		}
	}

	return snap, nil
}
