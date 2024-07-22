package user

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type UserHandlerImpl struct {
	store  types.UserStore
	logger *slog.Logger
}

func NewUserHandler(store types.UserStore, logger *slog.Logger) *UserHandlerImpl {
	return &UserHandlerImpl{
		store:  store,
		logger: logger,
	}
}
