package ec_point

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

var curve = elliptic.P521()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// BasePointGGet G-generator receiving
func (p *ECPoint) BasePointGGet() ECPoint {
	return ECPoint{
		X: curve.Params().Gx,
		Y: curve.Params().Gy,
	}
}

// ECPointGen ECPoint creation with pre-defined parameters
func (p *ECPoint) ECPointGen(x, y *big.Int) ECPoint {
	return ECPoint{
		X: x,
		Y: y,
	}
}

// IsOnCurveCheck P ∈ CURVE?
func (p *ECPoint) IsOnCurveCheck(a ECPoint) bool {
	return curve.IsOnCurve(a.X, a.Y)
}

// AddECPoints P + Q
func (p *ECPoint) AddECPoints(a, b ECPoint) ECPoint {
	return p.ECPointGen(curve.Add(a.X, a.Y, b.X, b.Y))
}

// DoubleECPoints 2P
func (p *ECPoint) DoubleECPoints(a ECPoint) ECPoint {
	return p.ECPointGen(curve.Double(a.X, a.Y))
}

// ScalarMult k * P
func (p *ECPoint) ScalarMult(a ECPoint, k big.Int) ECPoint {
	return p.ECPointGen(curve.ScalarMult(a.X, a.Y, k.Bytes()))
}

// ECPointToString Convert point to string
func (p *ECPoint) ECPointToString(point ECPoint) (s string) {
	return fmt.Sprintf("(%v; %v)", point.X, point.Y)
}

// PrintECPoint Print point
func (p *ECPoint) PrintECPoint(point ECPoint) {
	println(p.ECPointToString(point))
}
