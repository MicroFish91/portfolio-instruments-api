package account

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresAccountStore) CreateAccount(ctx context.Context, a *types.Account) error {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := s.db.Exec(
		c,
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
