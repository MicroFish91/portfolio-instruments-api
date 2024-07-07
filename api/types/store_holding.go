package types

type HoldingStore interface {
	CreateHolding(*Holding) error
	GetHoldings(int) (*[]Holding, error)
	GetHoldingById(int) (*Holding, error)
}
