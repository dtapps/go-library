package godecimal

import (
	"math/big"
)

type Decimal struct {
	floatValue *big.Float
}

func New() Decimal {
	return Decimal{
		floatValue: new(big.Float),
	}
}
