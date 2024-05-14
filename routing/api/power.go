package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func PowerChart(c *fiber.Ctx) error {
	days := gjson.Get(string(c.Body()), "days").Int()
	month := gjson.Get(string(c.Body()), "month").Int()
	ctx, client := es.GetESInstance()
	var powers []types.DailyPower
	if days != 0 {
		powers = es.QueryMinerDailySum(ctx, client, int(days))
		return c.JSON(pkg.SuccessResponse(powers))
	} else if month != 0 {
		powers = es.QueryMinerMonthSum(ctx, client, int(month))
		return c.JSON(pkg.SuccessResponse(powers))
	}
	return c.JSON(pkg.SuccessResponse(powers))
}
