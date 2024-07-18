package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GetSnapshotById(ctx context.Context, snapshotId, userId int) (*types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM snapshots
		WHERE user_id = $1
		AND snap_id = $2`,
		userId, snapshotId,
	)

	snapshot, err := s.parseRowIntoSnapshot(row)

	if err != nil {
		return nil, err
	}
	return snapshot, nil
}

func (s *PostgresSnapshotStore) parseRowIntoSnapshot(row pgx.Row) (*types.Snapshot, error) {
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
