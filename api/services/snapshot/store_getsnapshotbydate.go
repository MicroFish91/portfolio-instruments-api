package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) GetSnapshotByDate(ctx context.Context, snapshotDate string, userId int) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select 
				* 
			from 
				snapshots
			where 
				user_id = $1
				and snap_date = $2
		`,
		userId,
		snapshotDate,
	)

	return s.parseRowIntoSnapshot(row)
}
