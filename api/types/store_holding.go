package types

type GetHoldingsStoreOptions struct {
	Holding_ids    []int
	Ticker         string
	Asset_category AssetCategory
	Is_deprecated  string
}

type HoldingStore interface {
	CreateHolding(*Holding) error
	GetHoldings(int, *GetHoldingsStoreOptions) (*[]Holding, error)
	GetHoldingById(int) (*Holding, error)
}
