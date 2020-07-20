package common

import "github.com/shopspring/decimal"

func MustParseFloatStringFloat64(val string) float64 {
	quantity, err := decimal.NewFromString(val)
	if err != nil {
		panic(err)
	}
	quantityFloat, _ := quantity.Float64()
	return quantityFloat
}
