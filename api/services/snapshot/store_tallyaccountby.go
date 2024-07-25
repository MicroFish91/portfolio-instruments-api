package snapshot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) TallyAccountBy(ctx context.Context, snapId, userId int, options *types.GetTallyByAccountStoreOptions) (*types.AccountsGrouped, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if options == nil {
		return nil, errors.New("required to designate a tally_by options parameter")
	}

	var field string
	if options.Tally_by == types.BY_ACCOUNT_NAME {
		field = "name"
	} else if options.Tally_by == types.BY_ACCOUNT_INSTITUTION {
		field = "institution"
	}

	rows, err := s.db.Query(
		c,
		fmt.Sprintf(
			`SELECT accounts.%s, SUM(snapshots_values.total) AS total
			FROM snapshots_values
			INNER JOIN accounts
			ON snapshots_values.account_id = accounts.account_id
			WHERE snapshots_values.user_id = $1
			AND snapshots_values.snap_id = $2 
			AND is_deprecated is false
			GROUP BY accounts.%s`,
			field, field,
		),
		userId, snapId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accountsGrouped, err := s.parseRowsIntoAccountsGrouped(rows)

	if err != nil {
		return nil, err
	}
	return accountsGrouped, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoAccountsGrouped(rows pgx.Rows) (*types.AccountsGrouped, error) {
	type AccountGrouped struct {
		Field string
		Total float64
	}

	var ags types.AccountsGrouped
	for rows.Next() {
		var ag AccountGrouped
		err := rows.Scan(
			&ag.Field,
			&ag.Total,
		)

		if err != nil {
			return nil, err
		}

		ags.Fields = append(ags.Fields, ag.Field)
		ags.Total = append(ags.Total, ag.Total)
	}

	return &ags, nil
}
