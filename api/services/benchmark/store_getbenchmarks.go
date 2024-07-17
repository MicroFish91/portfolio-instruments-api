package benchmark

import (
	"context"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresBenchmarkStore) GetBenchmarks(ctx context.Context, userId int, options *types.GetBenchmarksStoreOptions) (*[]types.Benchmark, *types.PaginationMetadata, error) {
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
	pgxb.AddQuery("FROM benchmarks")
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})

	if options.Name != "" {
		pgxb.AddQueryWithPositionals("AND name ~* $x", []any{options.Name})
	}

	if options.Is_deprecated != "" {
		pgxb.AddQueryWithPositionals("AND is_deprecated = $x", []any{options.Is_deprecated})
	}

	if options.Benchmark_ids != nil && len(options.Benchmark_ids) > 0 {
		pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND benchmark_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.Benchmark_ids))),
			utils.IntSliceToAny(options.Benchmark_ids),
		)
	}

	pgxb.AddQuery("ORDER BY created_at ASC")
	pgxb.AddQuery(fmt.Sprintf("LIMIT %d OFFSET %d", pageSize, (currentPage-1)*pageSize))

	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.Query(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var total_items int
	var benchmarks []types.Benchmark
	for rows.Next() {
		var b types.Benchmark
		err = rows.Scan(
			&b.Benchmark_id,
			&b.Name,
			&b.Description,
			&b.Asset_allocation,
			&b.Std_dev_pct,
			&b.Real_return_pct,
			&b.Drawdown_yrs,
			&b.Is_deprecated,
			&b.User_id,
			&b.Created_at,
			&b.Updated_at,
			&total_items,
		)

		if err != nil {
			return nil, nil, err
		}
		benchmarks = append(benchmarks, b)
	}

	return &benchmarks, &types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  total_items,
	}, nil
}
