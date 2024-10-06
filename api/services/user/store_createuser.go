package user

import (
	"context"
	"errors"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) CreateUser(ctx context.Context, u *types.User) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if u == nil {
		return types.User{}, errors.New("internal error: user struct cannot be nil, valid user data is required")
	}

	row := s.db.QueryRow(
		c,
		`INSERT INTO users 
		(email, enc_password) 
		VALUES ($1, $2)
		RETURNING *`,
		strings.ToLower(u.Email), u.Enc_password,
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
