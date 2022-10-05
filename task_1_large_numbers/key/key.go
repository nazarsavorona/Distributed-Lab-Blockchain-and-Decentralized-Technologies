package key

import (
	"crypto/rand"
	"math/big"
)

func keySpace(bitsCount int64) *big.Int {
	bits := big.NewInt(bitsCount)
	two := big.NewInt(2)

	return two.Exp(two, bits, nil)
}

func generateKey(bitsCount int64) (*big.Int, error) {
	max := new(big.Int)
	max = keySpace(bitsCount).Sub(max, big.NewInt(1))

	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		return nil, err
	}

	return n, nil
}
