package auth

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type AuthHandlerImpl struct {
	store               types.UserStore
	logger              *slog.Logger
	jwtSecret           string
	requireVerification bool
}

func NewAuthHandler(store types.UserStore, logger *slog.Logger, jwtSecret string, requireVerification bool) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		store:               store,
		logger:              logger,
		jwtSecret:           jwtSecret,
		requireVerification: requireVerification,
	}
}
