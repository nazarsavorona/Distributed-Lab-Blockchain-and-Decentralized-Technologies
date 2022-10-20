package endian

import (
	"math/big"
)

type BigEndian []byte

func (bE *BigEndian) String() string {
	integerRepresentation := big.NewInt(0)
	power := big.NewInt(int64(len(*bE) - 1))
	one := big.NewInt(1)
	base := big.NewInt(256)

	for _, currentByte := range *bE {
		exp := big.NewInt(1).Exp(base, power, nil)
		integerRepresentation = integerRepresentation.Add(integerRepresentation, exp.Mul(exp, big.NewInt(int64(currentByte))))
		power.Sub(power, one)
	}

	return integerRepresentation.String()
}
