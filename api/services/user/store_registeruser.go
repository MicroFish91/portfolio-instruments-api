package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) RegisterUser(ctx context.Context, u *types.User) error {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := s.db.Exec(
		c,
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
