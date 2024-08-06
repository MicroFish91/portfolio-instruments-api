package account

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) UpdateAccount(ctx context.Context, account types.Account) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			update
				accounts
			set
				name = $1,
				description = $2,
				tax_shelter = $3,
				institution = $4,
				is_deprecated = $5,
				updated_at = now()
			where
				account_id = $6
				and user_id = $7
			returning
				*
		`,
		account.Name,
		account.Description,
		account.Tax_shelter,
		account.Institution,
		account.Is_deprecated,
		account.Account_id,
		account.User_id,
	)

	a, err := s.parseRowIntoAccount(row)
	if err != nil {
		return types.Account{}, err
	}
	return a, nil
}