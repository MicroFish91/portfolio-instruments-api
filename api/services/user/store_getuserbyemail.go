package user

import (
	"context"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				*
			from 
				users 
			where 
				email = $1`,
		strings.ToLower(email),
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (s *PostgresUserStore) parseRowIntoUser(row pgx.Row) (types.User, error) {
	var u types.User
	err := row.Scan(
		&u.User_id,
		&u.Email,
		&u.Enc_password,
		&u.User_role,
		&u.Last_logged_in,
		&u.Verified,
		&u.Created_at,
		&u.Updated_at,
	)

	if err != nil {
		return types.User{}, err
	}
	return u, nil
}
