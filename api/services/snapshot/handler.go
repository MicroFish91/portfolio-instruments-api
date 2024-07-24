package snapshot

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type SnapshotHandlerImpl struct {
	benchmarkStore types.BenchmarkStore
	accountStore   types.AccountStore
	holdingStore   types.HoldingStore
	snapshotStore  types.SnapshotStore
	logger         *slog.Logger
}

func NewSnapshotHandler(
	benchmarkStore types.BenchmarkStore,
	accountStore types.AccountStore,
	holdingStore types.HoldingStore,
	snapshotStore types.SnapshotStore,
	logger *slog.Logger,
) *SnapshotHandlerImpl {

	return &SnapshotHandlerImpl{
		benchmarkStore: benchmarkStore,
		accountStore:   accountStore,
		holdingStore:   holdingStore,
		snapshotStore:  snapshotStore,
		logger:         logger,
	}
}
