package pkg

import (
	"fmt"
	"math"
	"math/big"
)

func DivisionPowerOfTen(num float64, power int) float64 {
	return num / math.Pow(10, float64(power))
}

// DivisionBigPowerOfTen divides a big integer represented as a string by 10^power and returns the result as a float64
func DivisionBigPowerOfTen(numStr string, power int) (float64, error) {
	// 将字符串解析为 big.Int
	num := new(big.Int)
	_, success := num.SetString(numStr, 10)
	if !success {
		return 0, fmt.Errorf("invalid number string: %s", numStr)
	}

	// 计算10的power次方
	ten := big.NewInt(10)
	powerBigInt := new(big.Int).Exp(ten, big.NewInt(int64(power)), nil)

	// 将大整数除以10的power次方
	result := new(big.Float).Quo(new(big.Float).SetInt(num), new(big.Float).SetInt(powerBigInt))

	// 将结果转换为 float64
	floatResult, _ := result.Float64()

	return floatResult, nil
}
