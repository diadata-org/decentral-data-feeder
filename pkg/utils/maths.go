package utils

import (
	"math"
	"math/big"
)

func ScaleFloat(f float64, decimals int) *big.Int {
	fBig := big.NewFloat(f)
	scaling := big.NewFloat(math.Pow10(decimals))
	priceScaled := new(big.Float).Mul(fBig, scaling)
	valueUSDInt := new(big.Int)
	priceScaled.Int(valueUSDInt)
	return valueUSDInt
}

func ScaleInt(i int64, decimals int) *big.Int {
	fBig := new(big.Float).SetInt(big.NewInt(i))
	scaling := big.NewFloat(math.Pow10(decimals))
	priceScaled := new(big.Float).Mul(fBig, scaling)
	valueUSDInt := new(big.Int)
	priceScaled.Int(valueUSDInt)
	return valueUSDInt
}
