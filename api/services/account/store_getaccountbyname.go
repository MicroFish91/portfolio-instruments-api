package account

import (
	"context"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) GetAccountByName(ctx context.Context, name string, userId int) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	namePattern := fmt.Sprintf("^%s$", name)

	row := s.db.QueryRow(
		c,
		`SELECT * FROM accounts
		WHERE user_id = $1
		AND name ~* $2
		AND is_deprecated = false`,
		userId, namePattern,
	)

	account, err := s.parseRowIntoAccount(row)

	if err != nil {
		return types.Account{}, err
	}
	return account, nil
}
