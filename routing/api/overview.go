package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
)

// @Tags Web
// @Summary [Overview] OverviewInfo
// @Accept json
// @Description OverviewInfo API
// @Param param body types.Example false "Request Params"
// @Success 200 {object} types.ExampleRes "Success Response"
// @Router /overview/info [post]

func OverviewInfo(c *fiber.Ctx) error {
	//fmt.Print("开始模拟")
	//var msgReq types.OverviewInfoRes
	//err := c.BodyParser(&msgReq)
	//if err != nil {
	//	return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	//}
	lastBlock, err := es.GetLastBlock()
	if err != nil {
		return err
	}
	lb := lastBlock.Header
	info, err := es.GetNetWorkInfo()
	if err != nil {
		return err
	}

	sum := es.QueryChipsPower()
	totalReward24 := es.QueryBlockReward24h()
	var aveOut24 float64
	if sum != 0 {
		aveOut24 = float64(totalReward24 / sum)
	}

	totalMsgs24 := es.QueryBlockChangeMsg24h()
	ex := types.OverviewInfoRes{
		//Height:           pkg.FakeInt(0, 100000),
		//LatestBlock:      pkg.FakeIntStr(10, 120),
		//TotalPower:       pkg.FakeIntStr(10000, 50000),
		//ActiveMiner:      pkg.FakeIntStr(1000, 5000),
		//BlockReward:      pkg.FakeFloat(0, 1, 3),
		//DayAveReward:     pkg.FakeFloat(0, 1, 2),
		//DayProduction:    pkg.FakeIntStr(10000, 100000),
		//DayMessage:       pkg.FakeIntStr(10000, 20000),
		//TotalAccount:     pkg.FakeIntStr(5000, 10000),
		//AveBlockInterval: pkg.FakeIntStr(10, 60),
		Height:           lb.Height,
		LatestBlock:      lb.TimestampNanosec,
		TotalPower:       sum,
		ActiveMiner:      info.NumActivePeers,
		BlockReward:      totalReward24,
		DayAveReward:     aveOut24,
		DayProduction:    pkg.FakeIntStr(10000, 100000),
		DayMessages:      totalMsgs24,
		TotalAccount:     info.PeerMaxCount,
		AveBlockInterval: pkg.FakeIntStr(28, 33),
	}
	return c.JSON(pkg.SuccessResponse(ex))
}
