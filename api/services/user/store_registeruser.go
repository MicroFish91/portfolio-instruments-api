package user

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) RegisterUser(ctx context.Context, u *types.User) error {
	_, err := s.db.Exec(
		ctx,
		`INSERT INTO users 
		(email, enc_password) 
		VALUES ($1, $2)`,
		u.Email, u.Enc_password,
	)
	if err != nil {
		return err
	}
	return nil
}
