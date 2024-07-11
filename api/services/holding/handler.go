package holding

import "github.com/MicroFish91/portfolio-instruments-api/api/types"

type HoldingHandlerImpl struct {
	store types.HoldingStore
}

func NewHoldingHandler(store types.HoldingStore) *HoldingHandlerImpl {
	return &HoldingHandlerImpl{
		store: store,
	}
}
