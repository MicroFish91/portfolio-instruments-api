package user

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) GetSettings(ctx context.Context, userId int) (*types.Settings, error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row := s.db.QueryRow(
		c,
		`SELECT * FROM settings
		WHERE user_id = $1`,
		userId,
	)

	settings, err := s.parseRowIntoSettings(row)

	if err != nil {
		return nil, err
	}
	return settings, nil
}
