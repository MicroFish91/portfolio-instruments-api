package snapshotvalue

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotValueHandlerImpl struct {
	accountStore       types.AccountStore
	holdingStore       types.HoldingStore
	snapshotStore      types.SnapshotStore
	snapshotvalueStore types.SnapshotValueStore
	logger             *slog.Logger
}

func NewSnapshotValueHandler(
	accountStore types.AccountStore,
	holdingStore types.HoldingStore,
	snapshotStore types.SnapshotStore,
	snapshotvalueStore types.SnapshotValueStore,
	logger *slog.Logger,
) *SnapshotValueHandlerImpl {

	return &SnapshotValueHandlerImpl{
		accountStore:       accountStore,
		holdingStore:       holdingStore,
		snapshotStore:      snapshotStore,
		snapshotvalueStore: snapshotvalueStore,
		logger:             logger,
	}
}
