package account

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) CreateAccount(a *types.Account) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO accounts
		(name, description, tax_shelter, institution, is_deprecated, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		a.Name, a.Description, a.Tax_shelter, a.Institution, a.Is_deprecated, a.User_id,
	)
	if err != nil {
		return err
	}
	return nil
}
