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
	//powers := es.QueryMinerRange(ctx, client, days)
	powers := es.QueryMinerDailySum(ctx, client, int(days))
	return c.JSON(pkg.SuccessResponse(powers))
}
