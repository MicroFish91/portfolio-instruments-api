package user

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

const userColumns = `
	user_id,
	email,
	enc_password,
	user_role,
	last_logged_in,
	verified,
	created_at,
	updated_at
`

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

func (s *PostgresUserStore) parseRowsIntoUsers(rows pgx.Rows) ([]types.User, int, error) {
	var users []types.User
	var total_items int

	for rows.Next() {
		var u types.User
		err := rows.Scan(
			&u.User_id,
			&u.Email,
			&u.Enc_password,
			&u.User_role,
			&u.Last_logged_in,
			&u.Verified,
			&u.Created_at,
			&u.Updated_at,
			&total_items,
		)

		if err != nil {
			return []types.User{}, 0, nil
		}
		users = append(users, u)
	}
	return users, total_items, nil
}
