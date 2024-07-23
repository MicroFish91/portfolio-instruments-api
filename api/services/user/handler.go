package user

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type UserHandlerImpl struct {
	userStore      types.UserStore
	benchmarkStore types.BenchmarkStore
	logger         *slog.Logger
}

func NewUserHandler(userStore types.UserStore, benchmarkStore types.BenchmarkStore, logger *slog.Logger) *UserHandlerImpl {
	return &UserHandlerImpl{
		userStore:      userStore,
		benchmarkStore: benchmarkStore,
		logger:         logger,
	}
}
