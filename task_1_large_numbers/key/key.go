package key

import (
	"math/big"
)

func keySpace(bitsCount int64) *big.Int {
	bits := big.NewInt(bitsCount)
	two := big.NewInt(2)

	return two.Exp(two, bits, nil)
}
