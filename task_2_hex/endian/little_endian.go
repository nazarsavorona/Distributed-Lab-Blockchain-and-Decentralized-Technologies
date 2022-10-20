package endian

import "math/big"

type LittleEndian []byte

func (lE *LittleEndian) String() string {
	integerRepresentation := big.NewInt(0)
	power := big.NewInt(int64(len(*lE) - 1))
	one := big.NewInt(1)
	base := big.NewInt(256)

	for _, currentByte := range *lE {
		exp := big.NewInt(1).Exp(base, power, nil)
		integerRepresentation = integerRepresentation.Add(integerRepresentation, exp.Mul(exp, big.NewInt(int64(currentByte))))
		power.Sub(power, one)
	}

	return integerRepresentation.String()
}
