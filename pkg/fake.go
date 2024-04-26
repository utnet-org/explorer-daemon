package pkg

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
)

func FakeIntStr(a, b int) string {
	str := strconv.Itoa(gofakeit.Number(a, b))
	return str
}

func FakeInt(a, b int) int64 {
	return int64(gofakeit.Number(a, b))
}

func FakeFloat(a, b float64, place int) string {
	f := gofakeit.Float64Range(a, b)
	formattedFloat := fmt.Sprintf(fmt.Sprintf("%%.%df", place), f)
	return formattedFloat
}
