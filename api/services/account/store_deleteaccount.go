package account

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) DeleteAccount(ctx context.Context, userId, accountId int) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			delete from
				accounts
			where
				account_id = $1
				and user_id = $2
			returning 
				*
		`,
		accountId,
		userId,
	)

	account, err := s.parseRowIntoAccount(row)
	if err != nil {
		return types.Account{}, err
	}
	return account, nil
}
