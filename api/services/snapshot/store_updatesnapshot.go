package snapshot

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) UpdateSnapshot(ctx context.Context, snap *types.Snapshot) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if snap == nil {
		return types.Snapshot{}, errors.New("service error: snapshot struct cannot be nil, valid snapshot data is required")
	}

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			update
				snapshots
			set
				description = $1,
				snap_date = $2,
				total = $3,
				weighted_er_pct = $4,
				rebalance_threshold_pct = $5,
				value_order = $6,
				benchmark_id = $7,
				updated_at = now()
			where
				snap_id = $8
				and user_id = $9
			returning
				%s
		`, snapshotColumns),
		snap.Description,
		snap.Snap_date,
		snap.Total,
		snap.Weighted_er_pct,
		snap.Rebalance_threshold_pct,
		snap.Value_order,
		snap.Benchmark_id,
		snap.Snap_id,
		snap.User_id,
	)

	snapshot, err := s.parseRowIntoSnapshot(row)
	if err != nil {
		return types.Snapshot{}, err
	}
	return snapshot, nil
}
