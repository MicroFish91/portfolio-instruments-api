package snapshot

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresSnapshotStore) GroupByMaturationDate(ctx context.Context, userId, snapId int, options types.GetGroupByMaturationDateStoreOptions) ([]types.MaturationDateResource, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery(
		`
			select
				a.name,
				h.name,
				h.asset_category,
				h.interest_rate_pct,
				h.maturation_date,
				sv.total,
				sv.skip_rebalance
			from 
				snapshots_values sv
			inner join
				holdings h on h.holding_id = sv.holding_id
			inner join
				accounts a on a.account_id = sv.account_id 
			where 
		`,
	)

	pgxb.AddQueryWithPositionals("sv.user_id = $x", []any{userId})
	pgxb.AddQueryWithPositionals("and sv.snap_id = $x", []any{snapId})
	pgxb.AddQuery("and h.maturation_date != ''")

	if options.Maturation_start != "" {
		pgxb.AddQueryWithPositionals("and to_date(h.maturation_date, 'MM/DD/YYYY') >= to_date($x, 'MM/DD/YYYY')", []any{options.Maturation_start})
	}

	if options.Maturation_end != "" {
		pgxb.AddQueryWithPositionals("and to_date(h.maturation_date, 'MM/DD/YYYY') <= to_date($x, 'MM/DD/YYYY')", []any{options.Maturation_end})
	}

	pgxb.AddQuery(
		`
			order by 
				h.maturation_date desc,
				h.interest_rate_pct desc,
				sv.total desc
		`,
	)

	rows, err := s.db.Query(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resources, err := s.parseRowsIntoMaturationDateResources(rows)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (s *PostgresSnapshotStore) parseRowsIntoMaturationDateResources(rows pgx.Rows) ([]types.MaturationDateResource, error) {
	var resources []types.MaturationDateResource
	for rows.Next() {
		var r types.MaturationDateResource
		err := rows.Scan(
			&r.Account_name,
			&r.Holding_name,
			&r.Asset_category,
			&r.Interest_rate_pct,
			&r.Maturation_date,
			&r.Total,
			&r.Skip_rebalance,
		)

		if err != nil {
			return nil, err
		}
		resources = append(resources, r)
	}
	return resources, nil
}
