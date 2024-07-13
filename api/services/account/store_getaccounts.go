package account

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresAccountStore) GetAccounts(userId int, options *types.GetAccountsStoreOptions) (*[]types.Account, *types.PaginationMetadata, error) {
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
	pgxb.AddQuery("FROM accounts")
	err := pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})

	if options.Institution != "" {
		err = pgxb.AddQueryWithPositionals("AND institution ~* $x", []any{options.Institution})
	}
	if options.TaxShelter != "" {
		err = pgxb.AddQueryWithPositionals("AND tax_shelter = $x", []any{options.TaxShelter})
	}
	if options.Is_deprecated != "" {
		err = pgxb.AddQueryWithPositionals("AND is_deprecated = $x", []any{options.Is_deprecated})
	}

	if options.AccountIds != nil && len(options.AccountIds) > 0 {
		err = pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND account_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.AccountIds))),
			utils.IntSliceToAny(options.AccountIds),
		)
	}

	pgxb.AddQuery("ORDER BY created_at ASC")
	pgxb.AddQueryWithPositionals("LIMIT $x OFFSET $x", []any{pageSize, (currentPage - 1) * pageSize})

	if err != nil {
		return nil, nil, fmt.Errorf("error formatting sql query using query builder: %s", err.Error())
	}

	rows, err := s.db.Query(context.Background(), pgxb.Query, pgxb.QueryParams...)
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var total_items int
	var accounts []types.Account

	for rows.Next() {
		var a types.Account
		err := rows.Scan(
			&a.Account_id,
			&a.Name,
			&a.Description,
			&a.Tax_shelter,
			&a.Institution,
			&a.Is_deprecated,
			&a.User_id,
			&a.Created_at,
			&a.Updated_at,
			&total_items,
		)

		if err != nil {
			return nil, nil, err
		}
		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return &accounts, &types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  total_items,
	}, nil
}
