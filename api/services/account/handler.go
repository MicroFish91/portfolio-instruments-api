package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type AccountHandlerImpl struct {
	store types.AccountStore
}

func NewAccountHandler(store types.AccountStore) *AccountHandlerImpl {
	return &AccountHandlerImpl{
		store: store,
	}
}
