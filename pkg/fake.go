package pkg

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
)

func FakeInt(a, b int) string {
	str := strconv.Itoa(gofakeit.Number(a, b))
	return str
}

func FakeFloat(a, b float64, place int) string {
	f := gofakeit.Float64Range(a, b)
	formattedFloat := fmt.Sprintf(fmt.Sprintf("%%.%df", place), f)
	return formattedFloat
}
