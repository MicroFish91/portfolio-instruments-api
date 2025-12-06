package snapshotvalue

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

const snapshotValueColumns = `
	snap_val_id,
	snap_id,
	account_id,
	holding_id,
	total,
	skip_rebalance,
	user_id,
	created_at,
	updated_at
`

func (s *PostgresSnapshotValueStore) GetColumns() string {
	return snapshotValueColumns
}

func (s *PostgresSnapshotValueStore) parseRowIntoSnapshotValue(row pgx.Row) (types.SnapshotValue, error) {
	var sv types.SnapshotValue
	err := row.Scan(
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
		return types.SnapshotValue{}, err
	}
	return sv, nil
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
