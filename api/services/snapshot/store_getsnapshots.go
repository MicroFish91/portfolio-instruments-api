package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GetSnapshots(ctx context.Context, userId int) (*[]types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.Query(
		c,
		`SELECT * FROM snapshots
		WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	snapshots, err := s.parseRowsIntoSnapshots(rows)
	if err != nil {
		return nil, err
	}

	return snapshots, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoSnapshots(rows pgx.Rows) (*[]types.Snapshot, error) {
	var snapshots []types.Snapshot
	for rows.Next() {
		var s types.Snapshot
		err := rows.Scan(
			&s.Snap_id,
			&s.Description,
			&s.Snap_date,
			&s.Total,
			&s.User_id,
			&s.Created_at,
			&s.Updated_at,
		)

		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, s)
	}

	return &snapshots, nil
}
