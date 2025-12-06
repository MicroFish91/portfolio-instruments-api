package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) CreateUser(ctx context.Context, u *types.User) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	if u == nil {
		return types.User{}, errors.New("service error: user struct cannot be nil, valid user data is required")
	}

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			insert into users 
				(email, enc_password) 
				values ($1, $2)
			returning 
				%s
		`, userColumns),
		strings.ToLower(u.Email), u.Enc_password,
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
