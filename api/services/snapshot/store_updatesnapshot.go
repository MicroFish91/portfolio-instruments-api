package snapshot

import (
	"context"
	"errors"

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
		`
			update
				snapshots
			set
				description = $1,
				snap_date = $2,
				total = $3,
				weighted_er_pct = $4,
				benchmark_id = $5,
				updated_at = now()
			where
				snap_id = $6
				and user_id = $7
			returning
				*
		`,
		snap.Description,
		snap.Snap_date,
		snap.Total,
		snap.Weighted_er_pct,
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
