package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// @Tags Web
// @Summary [Overview] OverviewInfo
// @Accept json
// @Description OverviewInfo API
// @Param param body types.Example false "Request Params"
// @Success 200 {object} types.ExampleRes "Success Response"
// @Router /overview/info [post]

func OverviewInfo(c *fiber.Ctx) error {
	ctx, client := es.GetESInstance()
	lastBlock, err := es.GetLastBlock(ctx, client)
	if err != nil {
		log.Errorf("[OverviewInfo] GetLastBlock error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	lbh := lastBlock.Header
	info, err := es.GetNetworkInfo(ctx, client)
	if err != nil {
		log.Errorf("[OverviewInfo] GetNetworkInfo error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	val, err := es.QueryValidator(ctx, client)
	if err != nil {
		log.Errorf("[OverviewInfo] QueryValidator error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	//sum := es.QueryChipsPower(ctx, client)
	//totalReward24 := es.QueryBlockReward24h()

	blockReward := es.QueryRewordDiff(ctx, client)
	blockSupply24 := es.QuerySupplyDiff24h(ctx, client)

	miners, err := es.QueryMiner(ctx, client)
	if err != nil {
		return err
	}
	totalPower := pkg.DivisionPowerOfTen(float64(miners.TotalPower), 12)
	var aveOut24 float64
	if totalPower != 0 {
		formattedValue := fmt.Sprintf("%.6f", totalPower/blockSupply24)
		aveOut24, err = strconv.ParseFloat(formattedValue, 64)
		if err != nil {
			fmt.Println("Error converting to float:", err)
		}
	}
	totalMsgs24 := es.QueryBlockChangeMsg24h()
	ex := types.OverviewInfoRes{
		Height:           lbh.Height,
		LatestBlock:      strconv.Itoa(int(lbh.Timestamp)),
		TotalPower:       int64(totalPower),
		ActiveMiner:      int64(len(val.CurrentValidators)),
		BlockReward:      int64(blockReward),
		DayAveReward:     aveOut24,
		DayProduction:    int64(blockSupply24),
		DayMessages:      totalMsgs24,
		TotalAccount:     info.NumActivePeers,
		AveBlockInterval: pkg.FakeIntStr(28, 32),
	}
	return c.JSON(pkg.SuccessResponse(ex))
}
