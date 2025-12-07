package account

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) GetAccountByName(ctx context.Context, name string, userId int) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	namePattern := fmt.Sprintf("^%s$", name)

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select
				%s
			from 
				accounts
			where 
				user_id = $1
				and name ~* $2
		`, accountColumns),
		userId,
		namePattern,
	)

	account, err := s.parseRowIntoAccount(row)

	if err != nil {
		return types.Account{}, err
	}
	return account, nil
}
