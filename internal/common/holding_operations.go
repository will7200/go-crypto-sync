package common

type HoldingOperation int

const (
	AddHolding = 1 << iota
	RemoveHolding
	UpdateHolding
	SkipZeroQuantity
)

func (ao HoldingOperation) Has(flag HoldingOperation) bool { return ao&flag != 0 }
