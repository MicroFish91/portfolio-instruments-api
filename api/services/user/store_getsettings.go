package user

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) GetSettings(ctx context.Context, userId int) (types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM settings
		WHERE user_id = $1`,
		userId,
	)

	settings, err := s.parseRowIntoSettings(row)

	if err != nil {
		return types.Settings{}, err
	}
	return settings, nil
}
