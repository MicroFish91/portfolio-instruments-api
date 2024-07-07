package types

type HoldingStore interface {
	CreateHolding(*Holding) error
}
