package holding

type HoldingHandlerImpl struct {
	store *PostgresHoldingStore
}

func NewHoldingHandler(store *PostgresHoldingStore) *HoldingHandlerImpl {
	return &HoldingHandlerImpl{
		store: store,
	}
}
