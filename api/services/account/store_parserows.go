package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

const accountColumns = `
	account_id,
	name,
	description,
	tax_shelter,
	institution,
	is_deprecated,
	user_id,
	created_at,
	updated_at
`

func (s *PostgresAccountStore) parseRowIntoAccount(row pgx.Row) (types.Account, error) {
	var a types.Account
	err := row.Scan(
		&a.Account_id,
		&a.Name,
		&a.Description,
		&a.Tax_shelter,
		&a.Institution,
		&a.Is_deprecated,
		&a.User_id,
		&a.Created_at,
		&a.Updated_at,
	)

	if err != nil {
		return types.Account{}, err
	}
	return a, nil
}

func (s *PostgresAccountStore) parseRowsIntoAccounts(rows pgx.Rows) ([]types.Account, int, error) {
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
			return nil, 0, err
		}
		accounts = append(accounts, a)
	}

	return accounts, total_items, nil
}
