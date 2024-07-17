package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) CreateSnapshot(ctx context.Context, snapshot *types.Snapshot) (*types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO snapshots
		(snap_date, user_id)
		VALUES ($1, $2)
		RETURNING *`,
		snapshot.Snap_date, snapshot.User_id,
	)

	var snap types.Snapshot
	err := row.Scan(
		&snap.Snap_id,
		&snap.Snap_date,
		&snap.Total,
		&snap.User_id,
		&snap.Created_at,
		&snap.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &snap, nil
}
