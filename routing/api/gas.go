package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GasChart(c *fiber.Ctx) error {
	days := gjson.Get(string(c.Body()), "days").Int()
	ctx, client := es.GetESInstance()
	gass := es.QueryGasRange(ctx, client, days)
	return c.JSON(pkg.SuccessResponse(gass))
}
