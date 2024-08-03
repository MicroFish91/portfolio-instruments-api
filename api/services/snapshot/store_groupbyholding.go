package snapshot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GroupByHolding(ctx context.Context, userId, snapId int, options types.GetGroupByHoldingStoreOptions) (types.ResourcesGrouped, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	if options.Group_by == "" {
		return types.ResourcesGrouped{}, errors.New("required to designate a group_by options parameter")
	}

	var field string
	if options.Group_by == types.BY_ASSET_CATEGORY {
		field = "asset_category"
	}

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery(
		fmt.Sprintf("SELECT holdings.%s, SUM(snapshots_values.total) AS total", field),
	)
	pgxb.AddQuery(
		`FROM snapshots_values
		INNER JOIN holdings
		ON snapshots_values.holding_id = holdings.holding_id`,
	)
	pgxb.AddQueryWithPositionals("WHERE snapshots_values.user_id = $x", []any{userId})
	pgxb.AddQueryWithPositionals("AND snapshots_values.snap_id = $x", []any{snapId})
	pgxb.AddQuery("AND is_deprecated is false")

	if options.Omit_skip_reb {
		pgxb.AddQuery("AND snapshots_values.skip_rebalance is false")
	}

	pgxb.AddQuery(
		fmt.Sprintf("GROUP BY holdings.%s", field),
	)

	// search by upper and lower maturation date?

	rows, err := s.db.Query(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	if err != nil {
		return types.ResourcesGrouped{}, err
	}
	defer rows.Close()

	holdingsGrouped, err := s.parseRowsIntoHoldingsGrouped(rows)

	if err != nil {
		return types.ResourcesGrouped{}, err
	}
	return holdingsGrouped, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoHoldingsGrouped(rows pgx.Rows) (types.ResourcesGrouped, error) {
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
			return types.ResourcesGrouped{}, err
		}

		hgs.Fields = append(hgs.Fields, hg.Field)
		hgs.Total = append(hgs.Total, hg.Total)
	}

	return hgs, nil
}
