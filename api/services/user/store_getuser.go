package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT user_id, email, enc_password, created_at, updated_at 
		FROM users 
		WHERE email = $1`,
		email,
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *PostgresUserStore) parseRowIntoUser(row pgx.Row) (*types.User, error) {
	var u types.User
	err := row.Scan(&u.User_id, &u.Email, &u.Enc_password, &u.Created_at, &u.Updated_at)

	if err != nil {
		return nil, err
	}
	return &u, nil
}
