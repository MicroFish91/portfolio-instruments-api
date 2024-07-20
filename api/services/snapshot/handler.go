package snapshot

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotHandlerImpl struct {
	accountStore  types.AccountStore
	holdingStore  types.HoldingStore
	snapshotStore types.SnapshotStore
	logger        *slog.Logger
}

func NewSnapshotHandler(
	accountStore types.AccountStore,
	holdingStore types.HoldingStore,
	snapshotStore types.SnapshotStore,
	logger *slog.Logger,
) *SnapshotHandlerImpl {

	return &SnapshotHandlerImpl{
		accountStore:  accountStore,
		holdingStore:  holdingStore,
		snapshotStore: snapshotStore,
		logger:        logger,
	}
}
