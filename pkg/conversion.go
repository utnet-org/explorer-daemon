package pkg

import "math"

func DivisionPowerOfTen(num int64, power int) float64 {
	return float64(num) / math.Pow(10, float64(power))
}
