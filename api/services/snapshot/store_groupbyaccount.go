package snapshot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GroupByAccount(ctx context.Context, snapId, userId int, options types.GetGroupByAccountStoreOptions) (types.ResourcesGrouped, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if options.Group_by == "" {
		return types.ResourcesGrouped{}, errors.New("required to designate a group_by options parameter")
	}

	var field string
	if options.Group_by == types.BY_ACCOUNT_NAME {
		field = "name"
	} else if options.Group_by == types.BY_ACCOUNT_INSTITUTION {
		field = "institution"
	} else {
		field = "tax_shelter"
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
		return types.ResourcesGrouped{}, err
	}
	defer rows.Close()

	accountsGrouped, err := s.parseRowsIntoAccountsGrouped(rows)

	if err != nil {
		return types.ResourcesGrouped{}, err
	}
	return accountsGrouped, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoAccountsGrouped(rows pgx.Rows) (types.ResourcesGrouped, error) {
	type AccountGrouped struct {
		Field string
		Total float64
	}

	var ags types.ResourcesGrouped
	for rows.Next() {
		var ag AccountGrouped
		err := rows.Scan(
			&ag.Field,
			&ag.Total,
		)

		if err != nil {
			return types.ResourcesGrouped{}, err
		}

		ags.Fields = append(ags.Fields, ag.Field)
		ags.Total = append(ags.Total, ag.Total)
	}

	return ags, nil
}
