package account

import (
	"context"
	"fmt"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresAccountStore) GetAccounts(userId int, accountIds []int) (*[]types.Account, error) {
	var rows pgx.Rows
	var err error
	if len(accountIds) == 0 {
		rows, err = s.getAllAccountsRows(userId)
	} else {
		rows, err = s.getSomeAccountsRows(userId, accountIds)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []types.Account
	for rows.Next() {
		var a types.Account
		if err := rows.Scan(&a.Account_id, &a.Name, &a.Description, &a.Tax_shelter, &a.Institution, &a.Is_closed, &a.User_id, &a.Created_at, &a.Updated_at); err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (s *PostgresAccountStore) getAllAccountsRows(userId int) (pgx.Rows, error) {
	return s.db.Query(
		context.Background(),
		`SELECT *
		FROM accounts
		WHERE user_id = $1`,
		userId,
	)
}

func (s *PostgresAccountStore) getSomeAccountsRows(userId int, accountIds []int) (pgx.Rows, error) {
	inParams := make([]string, len(accountIds))

	queryArgs := make([]interface{}, len(accountIds)+1)
	queryArgs[0] = userId

	for i := 0; i < len(accountIds); i += 1 {
		p := fmt.Sprintf("$%d", i+2)
		inParams[i] = p
		queryArgs[i+1] = accountIds[i]
	}

	query := fmt.Sprintf(
		`SELECT *
		FROM accounts
		WHERE user_id = $1
		AND account_id IN (%s)`,
		strings.Join(inParams, ", "),
	)

	return s.db.Query(context.Background(), query, queryArgs...)
}
