package user

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) UpdateVerification(ctx context.Context, userId int) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			update
				users
			set
				verified = true,
				updated_at = now()
			where
				user_id = $1
			returning
				%s
		`, userColumns),
		userId,
	)

	user, err := s.parseRowIntoUser(row)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
