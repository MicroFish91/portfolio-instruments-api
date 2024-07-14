package holding

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type HoldingHandlerImpl struct {
	store  types.HoldingStore
	logger *slog.Logger
}

func NewHoldingHandler(store types.HoldingStore, logger *slog.Logger) *HoldingHandlerImpl {
	return &HoldingHandlerImpl{
		store:  store,
		logger: logger,
	}
}
