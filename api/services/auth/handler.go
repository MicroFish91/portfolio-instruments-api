package auth

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type AuthHandlerImpl struct {
	store  types.UserStore
	logger *slog.Logger
}

func NewAuthHandler(store types.UserStore, logger *slog.Logger) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		store:  store,
		logger: logger,
	}
}
