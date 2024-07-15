package snapshot

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotHandlerImpl struct {
	store  types.SnapshotStore
	logger *slog.Logger
}

func NewSnapshotHandler(store types.SnapshotStore, logger *slog.Logger) *SnapshotHandlerImpl {
	return &SnapshotHandlerImpl{
		store:  store,
		logger: logger,
	}
}
