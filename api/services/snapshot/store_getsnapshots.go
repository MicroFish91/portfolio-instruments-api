package snapshot

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresSnapshotStore) GetSnapshots(ctx context.Context, userId int, options *types.GetSnapshotsStoreOptions) ([]types.Snapshot, types.PaginationMetadata, error) {
	if options == nil {
		options = &types.GetSnapshotsStoreOptions{
			Snap_ids:        []int{},
			Snap_date_lower: "",
			Snap_date_upper: "",
			Order_date_by:   "",
			Current_page:    1,
			Page_size:       50,
		}
	}

	currentPage := 1
	if options.Current_page > 1 {
		currentPage = options.Current_page
	}

	pageSize := 50
	if options.Page_size > 0 && options.Page_size < 50 {
		pageSize = options.Page_size
	}

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery(fmt.Sprintf("SELECT %s, COUNT(*) OVER() as total_items", snapshotColumns))
	pgxb.AddQuery("FROM snapshots")
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})

	if len(options.Snap_ids) > 0 {
		pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND snap_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.Snap_ids))),
			utils.ConvertIntSliceToAny(options.Snap_ids),
		)
	}

	if options.Snap_date_lower != "" {
		pgxb.AddQueryWithPositionals("AND TO_DATE(snap_date, 'MM/DD/YYYY') >= TO_DATE($x, 'MM/DD/YYYY')", []any{options.Snap_date_lower})
	}

	if options.Snap_date_upper != "" {
		pgxb.AddQueryWithPositionals("AND TO_DATE(snap_date, 'MM/DD/YYYY') <= TO_DATE($x, 'MM/DD/YYYY')", []any{options.Snap_date_upper})
	}

	if options.Order_date_by != "" {
		if options.Order_date_by == "ASC" {
			pgxb.AddQuery("ORDER BY TO_DATE(snap_date, 'MM/DD/YYYY') ASC")
		} else {
			pgxb.AddQuery("ORDER BY TO_DATE(snap_date, 'MM/DD/YYYY') DESC")
		}
	}

	pgxb.AddQueryWithPositionals("LIMIT $x OFFSET $x", []any{pageSize, (currentPage - 1) * pageSize})

	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	rows, err := s.db.Query(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	if err != nil {
		return nil, types.PaginationMetadata{}, err
	}

	snapshots, total_items, err := s.parseRowsIntoSnapshots(rows)
	if err != nil {
		return nil, types.PaginationMetadata{}, err
	}

	return snapshots, types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  total_items,
	}, nil
}
