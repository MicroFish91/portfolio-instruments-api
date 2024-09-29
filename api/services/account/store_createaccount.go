package account

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) CreateAccount(ctx context.Context, a types.Account) (types.Account, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO accounts
		(name, description, tax_shelter, institution, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		a.Name, a.Description, a.Tax_shelter, a.Institution, a.Is_deprecated, a.User_id,
	)

	account, err := s.parseRowIntoAccount(row)

	if err != nil {
		return types.Account{}, err
	}
	return account, nil
}
