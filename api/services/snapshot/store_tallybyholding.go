package snapshot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) TallyByHolding(ctx context.Context, userId, snapId int, options *types.GetTallyByHoldingStoreOptions) (*types.ResourcesGrouped, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if options == nil || options.Tally_by == "" {
		return nil, errors.New("required to designate a tally_by options parameter")
	}

	var field string
	if options.Tally_by == types.BY_ASSET_CATEGORY {
		field = "asset_category"
	}

	// search by upper and lower maturation date?

	rows, err := s.db.Query(
		c,
		fmt.Sprintf(
			`SELECT holdings.%s, SUM(snapshots_values.total) AS total
			FROM snapshots_values
			INNER JOIN holdings
			ON snapshots_values.holding_id = holdings.holding_id
			WHERE snapshots_values.user_id = $1
			AND snapshots_values.snap_id = $2
			AND is_deprecated is false
			GROUP BY holdings.%s`,
			field, field,
		),
		userId, snapId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	holdingsGrouped, err := s.parseRowsIntoHoldingsGrouped(rows)

	if err != nil {
		return nil, err
	}
	return holdingsGrouped, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoHoldingsGrouped(rows pgx.Rows) (*types.ResourcesGrouped, error) {
	type HoldingGroup struct {
		Field string
		Total float64
	}

	var hgs types.ResourcesGrouped
	for rows.Next() {
		var hg HoldingGroup
		err := rows.Scan(
			&hg.Field,
			&hg.Total,
		)

		if err != nil {
			return nil, err
		}

		hgs.Fields = append(hgs.Fields, hg.Field)
		hgs.Total = append(hgs.Total, hg.Total)
	}

	return &hgs, nil
}
