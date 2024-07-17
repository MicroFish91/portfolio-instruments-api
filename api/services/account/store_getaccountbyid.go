package account

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) GetAccountById(ctx context.Context, userId int, accountId int) (*types.Account, error) {
	row := s.db.QueryRow(
		ctx,
		`SELECT *
		FROM accounts
		WHERE user_id = $1 AND account_id = $2`,
		userId, accountId,
	)

	var a types.Account
	err := row.Scan(&a.Account_id, &a.Name, &a.Description, &a.Tax_shelter, &a.Institution, &a.Is_deprecated, &a.User_id, &a.Created_at, &a.Updated_at)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
