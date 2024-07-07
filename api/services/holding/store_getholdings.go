package holding

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresHoldingStore) GetHoldings(userId int, options *types.GetHoldingsStoreOptions) (*[]types.Holding, error) {
	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("SELECT * FROM holdings")
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

	rows, err := s.db.Query(context.Background(), pgxb.Query, pgxb.QueryParams...)
	if err != nil {
		return nil, err
	}

	var holdings []types.Holding
	for rows.Next() {
		var h types.Holding
		if err = rows.Scan(&h.Holding_id, &h.Name, &h.Ticker, &h.Asset_category, &h.Expense_ratio, &h.Is_deprecated, &h.User_id, &h.Created_at, &h.Updated_at); err != nil {
			return nil, err
		}
		holdings = append(holdings, h)
	}
	return &holdings, nil
}
