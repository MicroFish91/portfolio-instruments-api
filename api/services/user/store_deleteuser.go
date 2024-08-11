package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) DeleteUser(ctx context.Context, userId int) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err := s.db.Exec(
		c,
		`
			delete from
				settings
			where
				user_id = $1
		`,
		userId,
	)
	if err != nil {
		return types.User{}, err
	}

	row := s.db.QueryRow(
		c,
		`
			delete from
				users
			where
				user_id = $1
			returning
				*
		`,
		userId,
	)

	user, err := s.parseRowIntoUser(row)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
