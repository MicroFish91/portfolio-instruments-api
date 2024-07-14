package account

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type AccountHandlerImpl struct {
	store  types.AccountStore
	logger *slog.Logger
}

func NewAccountHandler(store types.AccountStore, logger *slog.Logger) *AccountHandlerImpl {
	return &AccountHandlerImpl{
		store:  store,
		logger: logger,
	}
}
