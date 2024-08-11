package user

import "github.com/MicroFish91/portfolio-instruments-api/api/types"

type GetUsersQuery struct {
	types.PaginationQuery
}

func (q GetUsersQuery) Validate() error {
	return q.PaginationQuery.Validate()
}
