package user

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/jackc/pgx/v5"
)

func (s *PostgresUserStore) GetUsers(ctx context.Context, options types.GetUsersStoreOptions) ([]types.User, types.PaginationMetadata, error) {
	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	currentPage := 1
	if options.Current_page > 1 {
		currentPage = options.Current_page
	}

	pageSize := 50
	if options.Page_size != 0 {
		pageSize = options.Page_size
	}

	rows, err := s.db.Query(
		c,
		`
			select
				*, 
				COUNT(*) OVER() as total_items
			from
				users
			order by
				last_logged_in ASC,
				updated_at ASC,
				created_at ASC
			limit
				$1
			offset
				$2	
		`,
		pageSize,
		(currentPage-1)*pageSize,
	)

	if err != nil {
		return []types.User{}, types.PaginationMetadata{}, err
	}
	defer rows.Close()

	users, totalItems, err := s.parseRowsIntoUsers(rows)
	if err != nil {
		return []types.User{}, types.PaginationMetadata{}, err
	}

	return users, types.PaginationMetadata{
		Current_page: currentPage,
		Page_size:    pageSize,
		Total_items:  totalItems,
	}, nil
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
