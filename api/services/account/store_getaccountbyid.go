package account

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) GetAccountById(ctx context.Context, userId int, accountId int) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select
				%s
			from
				accounts
			where 
				user_id = $1
				and account_id = $2
		`, accountColumns),
		userId,
		accountId,
	)

	account, err := s.parseRowIntoAccount(row)
	if err != nil {
		return types.Account{}, err
	}

	return account, nil
}
