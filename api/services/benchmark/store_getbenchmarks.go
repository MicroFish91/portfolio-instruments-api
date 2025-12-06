package benchmark

import (
	"context"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresBenchmarkStore) GetBenchmarks(ctx context.Context, userId int, options *types.GetBenchmarksStoreOptions) ([]types.Benchmark, types.PaginationMetadata, error) {
	if options == nil {
		options = &types.GetBenchmarksStoreOptions{
			Benchmark_ids: []int{},
			Name:          "",
			Is_deprecated: "",
			Current_page:  1,
			Page_size:     50,
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
	pgxb.AddQuery(fmt.Sprintf("SELECT %s, COUNT(*) OVER() as total_items", benchmarkColumns))
	pgxb.AddQuery("FROM benchmarks")
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})

	if options.Name != "" {
		pgxb.AddQueryWithPositionals("AND name ~* $x", []any{options.Name})
	}

	if options.Is_deprecated != "" {
		pgxb.AddQueryWithPositionals("AND is_deprecated = $x", []any{options.Is_deprecated})
	}

	if len(options.Benchmark_ids) > 0 {
		pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND benchmark_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.Benchmark_ids))),
			utils.ConvertIntSliceToAny(options.Benchmark_ids),
		)
	}

	pgxb.AddQuery("ORDER BY created_at ASC")
	pgxb.AddQuery(fmt.Sprintf("LIMIT %d OFFSET %d", pageSize, (currentPage-1)*pageSize))

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
	defer rows.Close()

	benchmarks, total_items, err := s.parseRowsIntoBenchmarks(rows)
	if err != nil {
		return nil, types.PaginationMetadata{}, err
	}
	if len(benchmarks) == 0 {
		return nil, types.PaginationMetadata{}, errors.New("no rows in result set")
	}

	return benchmarks, types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  total_items,
	}, nil
}
