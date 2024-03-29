package web

import (
	"explorer-daemon/pkg"
	"explorer-daemon/routing/types"
	"github.com/gofiber/fiber/v2"
)

// @Tags OverviewInfo
// @Summary OverviewInfo
// @Accept json
// @Description OverviewInfo API
// @Param param body false "Request Params"
// @Success 200 {object} types.OverviewInfo "Success Response"
// @Router /overview/info [post]
func OverviewInfo(c *fiber.Ctx) error {
	//fmt.Print("开始模拟")
	//var msgReq types.OverviewInfoRes
	//err := c.BodyParser(&msgReq)
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	//}
	ex := types.OverviewInfoRes{
		Height:           pkg.FakeInt(0, 100000),
		LatestBlock:      pkg.FakeInt(10, 120),
		TotalPower:       pkg.FakeInt(10000, 50000),
		ActiveMiner:      pkg.FakeInt(1000, 5000),
		BlockReward:      pkg.FakeFloat(0, 1, 3),
		DayAveReward:     pkg.FakeFloat(0, 1, 2),
		DayProduction:    pkg.FakeInt(10000, 100000),
		DayMessage:       pkg.FakeInt(10000, 20000),
		TotalAccount:     pkg.FakeInt(5000, 10000),
		AveBlockInterval: pkg.FakeInt(10, 60),
	}
	return c.JSON(pkg.SuccessResponse(ex))
}
