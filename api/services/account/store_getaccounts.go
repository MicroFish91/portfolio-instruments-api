package account

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
)

func (s *PostgresAccountStore) GetAccounts(userId int, options *types.GetAccountsStoreOptions) (*[]types.Account, error) {
	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("SELECT * FROM accounts")
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

	if len(options.AccountIds) > 0 {
		err = pgxb.AddQueryWithPositionals(
			fmt.Sprintf("AND account_id IN (%s)", querybuilder.FillWithEmptyPositionals(len(options.AccountIds))),
			utils.IntSliceToAny(options.AccountIds),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("error formatting sql query using query builder: %s", err.Error())
	}

	rows, err := s.db.Query(context.Background(), pgxb.Query, pgxb.QueryParams...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []types.Account
	for rows.Next() {
		var a types.Account
		if err := rows.Scan(&a.Account_id, &a.Name, &a.Description, &a.Tax_shelter, &a.Institution, &a.Is_deprecated, &a.User_id, &a.Created_at, &a.Updated_at); err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &accounts, nil
}
