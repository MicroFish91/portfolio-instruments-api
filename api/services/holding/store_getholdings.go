package holding

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresHoldingStore) GetHoldings(userId int, options *types.GetHoldingsStoreOptions) (*[]types.Holding, *types.PaginationMetadata, error) {
	currentPage := 1
	if options.Current_page > 1 {
		currentPage = options.Current_page
	}

	pageSize := 50
	if options.Page_size > 0 && options.Page_size < 50 {
		pageSize = options.Page_size
	}

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("SELECT *, COUNT(*) OVER() as total_items")
	pgxb.AddQuery("FROM holdings")
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})

	if options.Ticker != "" {
		pgxb.AddQueryWithPositionals("AND ticker ~* $x", []any{options.Ticker})
	}

	if options.Asset_category != "" {
		pgxb.AddQueryWithPositionals("AND asset_category = $x", []any{options.Asset_category})
	}

	if options.Is_deprecated != "" {
		pgxb.AddQueryWithPositionals("AND is_deprecated = $x", []any{options.Is_deprecated})
	}

	if options.Holding_ids != nil && len(options.Holding_ids) > 0 {
		pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND holding_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.Holding_ids))),
			utils.IntSliceToAny(options.Holding_ids),
		)
	}

	pgxb.AddQuery("ORDER BY created_at ASC")
	pgxb.AddQuery(fmt.Sprintf("LIMIT %d OFFSET %d", pageSize, (currentPage-1)*pageSize))

	rows, err := s.db.Query(context.Background(), pgxb.Query, pgxb.QueryParams...)
	if err != nil {
		return nil, nil, err
	}

	var total_items int
	var holdings []types.Holding

	for rows.Next() {
		var h types.Holding
		if err = rows.Scan(&h.Holding_id, &h.Name, &h.Ticker, &h.Asset_category, &h.Expense_ratio, &h.Is_deprecated, &h.User_id, &h.Created_at, &h.Updated_at, &total_items); err != nil {
			return nil, nil, err
		}
		holdings = append(holdings, h)
	}
	return &holdings, &types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  total_items,
	}, nil
}
