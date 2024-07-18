package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) RegisterUser(ctx context.Context, u *types.User) (*types.User, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`INSERT INTO users 
		(email, enc_password) 
		VALUES ($1, $2)
		RETURNING *`,
		u.Email, u.Enc_password,
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return nil, err
	}
	return user, nil
}
