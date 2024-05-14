package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func PowerChart(c *fiber.Ctx) error {
	days := gjson.Get(string(c.Body()), "days").Int()
	ctx, client := es.GetESInstance()
	powers := es.QueryMinerRange(ctx, client, days)
	finalPowers := make([]float64, 0)
	for _, p := range powers {
		finalP := pkg.DivisionPowerOfTen(p, 12)
		finalPowers = append(finalPowers, finalP)
	}
	return c.JSON(pkg.SuccessResponse(finalPowers))
}
