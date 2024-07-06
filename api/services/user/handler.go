package user

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type UserHandlerImpl struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandlerImpl {
	return &UserHandlerImpl{
		store: store,
	}
}
