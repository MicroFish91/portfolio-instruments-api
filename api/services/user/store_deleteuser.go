package user

import (
	"context"
	"fmt"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) DeleteUser(ctx context.Context, userId int) (types.User, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*40)
	defer cancel()

	foreignTables := []string{
		"snapshots_values",
		"snapshots",
		"holdings",
		"accounts",
		"settings",
		"benchmarks",
	}

	// Delete foreign tables
	for _, t := range foreignTables {
		_, err := s.db.Exec(
			c,
			fmt.Sprintf(
				`
					delete from
						%s
					where
						user_id = $1
				`,
				t,
			),
			userId,
		)
		if err != nil {
			return types.User{}, fmt.Errorf("failed to delete values from %s table: %s", t, err.Error())
		}
	}

	// Delete user
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
