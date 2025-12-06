package snapshotvalue

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotValueStore) CreateSnapshotValue(ctx context.Context, snapVals *types.SnapshotValue) (types.SnapshotValue, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if snapVals == nil {
		return types.SnapshotValue{}, errors.New("service error: snapshotvalue struct cannot be nil, valid snapshotvalue data is required")
	}

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			insert into snapshots_values
				(snap_id, account_id, holding_id, total, skip_rebalance, user_id)
				values ($1, $2, $3, $4, $5, $6)
			returning
				%s
		`, snapshotValueColumns),
		snapVals.Snap_id, snapVals.Account_id, snapVals.Holding_id, snapVals.Total, snapVals.Skip_rebalance, snapVals.User_id,
	)

	sv, err := s.parseRowIntoSnapshotValue(row)

	if err != nil {
		return types.SnapshotValue{}, err
	}
	return sv, nil
}
