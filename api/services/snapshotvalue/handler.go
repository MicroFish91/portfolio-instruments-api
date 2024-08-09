package snapshotvalue

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotValueHandlerImpl struct {
	store  types.SnapshotValueStore
	logger *slog.Logger
}

func NewSnapshotValueHandler(store types.SnapshotValueStore, logger *slog.Logger) *SnapshotValueHandlerImpl {
	return &SnapshotValueHandlerImpl{
		store:  store,
		logger: logger,
	}
}
