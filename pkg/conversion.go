package pkg

import "math"

func DivisionPowerOfTen(num float64, power int) float64 {
	return num / math.Pow(10, float64(power))
}
