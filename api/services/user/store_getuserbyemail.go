package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		fmt.Sprintf(`
			select
				%s
			from 
				users 
			where 
				email = $1
		`, userColumns),
		strings.ToLower(email),
	)

	user, err := s.parseRowIntoUser(row)

	if err != nil {
		return types.User{}, err
	}
	return user, nil
}
