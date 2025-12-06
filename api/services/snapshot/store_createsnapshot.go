package snapshot

import (
	"context"
	"errors"
	"fmt"

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
		fmt.Sprintf(`
			insert into snapshots
				(snap_date, description, user_id, rebalance_threshold_pct)
				values ($1, $2, $3, $4)
			returning
				%s
		`, snapshotColumns),
		snapshot.Snap_date, snapshot.Description, snapshot.User_id, snapshot.Rebalance_threshold_pct,
	)

	snap, err := s.parseRowIntoSnapshot(row)
	if err != nil {
		return types.Snapshot{}, err
	}

	// We skip adding the benchmark above because it could get defaulted to 0.  Instead, it's safer to have it show up as nil initially and then explicitly set it separately if necessary.
	if snapshot.Benchmark_id != 0 {
		row = s.db.QueryRow(
			c,
			fmt.Sprintf(`
				update
					snapshots
				set
					benchmark_id = $1
				where
					snap_id = $2
				returning
					%s 
			`, snapshotColumns),
			snapshot.Benchmark_id, snap.Snap_id,
		)

		snap, err = s.parseRowIntoSnapshot(row)
		if err != nil {
			return types.Snapshot{}, err
		}
	}

	return snap, nil
}
