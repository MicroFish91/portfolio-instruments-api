package snapshot

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotHandlerImpl struct {
	userStore      types.UserStore
	benchmarkStore types.BenchmarkStore
	accountStore   types.AccountStore
	holdingStore   types.HoldingStore
	snapshotStore  types.SnapshotStore
	logger         *slog.Logger
}

func NewSnapshotHandler(
	userStore types.UserStore,
	benchmarkStore types.BenchmarkStore,
	accountStore types.AccountStore,
	holdingStore types.HoldingStore,
	snapshotStore types.SnapshotStore,
	logger *slog.Logger,
) *SnapshotHandlerImpl {

	return &SnapshotHandlerImpl{
		userStore:      userStore,
		benchmarkStore: benchmarkStore,
		accountStore:   accountStore,
		holdingStore:   holdingStore,
		snapshotStore:  snapshotStore,
		logger:         logger,
	}
}
