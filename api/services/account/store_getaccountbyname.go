package account

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) GetAccountByName(ctx context.Context, name string, userId int) (*types.Account, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM accounts
		WHERE user_id = $1
		AND name = $2
		AND is_deprecated = false`,
		userId, name,
	)

	account, err := s.parseRowIntoAccount(row)

	if err != nil {
		return nil, err
	}
	return account, nil
}
