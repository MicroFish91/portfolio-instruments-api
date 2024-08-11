package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) GetUserById(ctx context.Context, userId int) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`
			select
				*
			from
				users
			where
				user_id = $1
		`,
		userId,
	)

	user, err := s.parseRowIntoUser(row)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
