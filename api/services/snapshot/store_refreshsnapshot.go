package snapshot

import "context"

func (s *PostgresSnapshotStore) RefreshSnapshotTotal(userId, snapshotId int) (float64, error) {
	// Use an aggregate function to sum row totals
	rows, err := s.db.Query(
		context.Background(),
		`SELECT SUM(total) AS snapshot_total
		FROM snapshots_values
		WHERE user_id = $1
		AND snap_id = $2`,
		userId, snapshotId,
	)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var snapshot_total float64
	for rows.Next() {
		if err := rows.Scan(&snapshot_total); err != nil {
			return 0, err
		}
	}

	// Update snapshots with the new total
	_, err = s.db.Exec(
		context.Background(),
		`UPDATE snapshots
		SET total = $1
		WHERE user_id = $2
		AND snap_id = $3`,
		snapshot_total, userId, snapshotId,
	)

	return snapshot_total, err
}
