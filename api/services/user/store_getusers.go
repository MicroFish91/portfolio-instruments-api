package user

import (
	"context"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresUserStore) GetUsers(ctx context.Context, options *types.GetUsersStoreOptions) ([]types.User, types.PaginationMetadata, error) {
	if options == nil {
		options = &types.GetUsersStoreOptions{
			Current_page: 1,
			Page_size:    50,
		}
	}

	if options.Current_page == 0 {
		options.Current_page = 1
	}

	if options.Page_size == 0 {
		options.Page_size = 50
	}

	c, cancel := context.WithTimeout(ctx, constants.TIMEOUT_MEDIUM)
	defer cancel()

	rows, err := s.db.Query(
		c,
		fmt.Sprintf(`
			select
				%s, 
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
		`, userColumns),
		options.Page_size,
		(options.Current_page-1)*options.Page_size,
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
		Current_page: options.Current_page,
		Page_size:    options.Page_size,
		Total_items:  totalItems,
	}, nil
}
