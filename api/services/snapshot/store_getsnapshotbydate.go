package snapshot

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) GetSnapshotByDate(ctx context.Context, snapshotDate string, userId int) (types.Snapshot, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select 
				%s 
			from 
				snapshots
			where 
				user_id = $1
				and snap_date = $2
		`, snapshotColumns),
		userId,
		snapshotDate,
	)

	return s.parseRowIntoSnapshot(row)
}
