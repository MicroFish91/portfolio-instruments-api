package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) CreateSnapshot(snapshot *types.Snapshot) (*types.Snapshot, error) {
	row := s.db.QueryRow(
		context.Background(),
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
