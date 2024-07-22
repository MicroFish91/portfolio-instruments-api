package user

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type UserHandlerImpl struct {
	userStore     types.UserStore
	settingsStore types.SettingsStore
	logger        *slog.Logger
}

func NewUserHandler(userStore types.UserStore, settingsStore types.SettingsStore, logger *slog.Logger) *UserHandlerImpl {
	return &UserHandlerImpl{
		userStore:     userStore,
		settingsStore: settingsStore,
		logger:        logger,
	}
}
