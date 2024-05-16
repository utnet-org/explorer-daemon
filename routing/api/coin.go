package api

import (
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
)

func CoinPrice(c *fiber.Ctx) error {
	price, change := remote.PriceHandler()
	res := types.CoinPriceResWeb{
		Price:  price,
		Amount: change,
	}
	return c.JSON(pkg.SuccessResponse(res))
}
