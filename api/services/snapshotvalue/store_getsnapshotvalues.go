package snapshotvalue

import (
	"context"
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotValueStore) GetSnapshotValues(ctx context.Context, snapId, userId int) ([]types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.Query(
		c,
		`
			select
				*
			from
				snapshots_values
			where
				snap_id = $1
				and user_id = $2
		`,
		snapId,
		userId,
	)

	if err != nil {
		return []types.SnapshotValue{}, err
	}
	defer rows.Close()

	snapshotValues, err := s.parseRowsIntoSnapshotValues(rows)
	if err != nil {
		return []types.SnapshotValue{}, err
	}
	if len(snapshotValues) == 0 {
		return []types.SnapshotValue{}, errors.New("snapshot_values for the given snapshot id not found")
	}

	return snapshotValues, nil
}

func (s *PostgresSnapshotValueStore) parseRowsIntoSnapshotValues(rows pgx.Rows) ([]types.SnapshotValue, error) {
	var svs []types.SnapshotValue
	for rows.Next() {
		var sv types.SnapshotValue
		err := rows.Scan(
			&sv.Snap_val_id,
			&sv.Snap_id,
			&sv.Account_id,
			&sv.Holding_id,
			&sv.Total,
			&sv.Skip_rebalance,
			&sv.User_id,
			&sv.Created_at,
			&sv.Updated_at,
		)

		if err != nil {
			return []types.SnapshotValue{}, err
		}
		svs = append(svs, sv)
	}

	return svs, nil
}
