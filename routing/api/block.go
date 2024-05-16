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
// @Summary [Block] BlockDetails
// @Accept json
// @Description BlockDetails API
// @Param param body types.BlockDetailsReq false "Request Params"
// @Success 200 {object} types.BlockDetailsRes "Success Response"
// @Router /block/details [post]
func BlockDetails(c *fiber.Ctx) error {
	var req types.BlockDetailsReq
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	}
	resWeb, err := BlockDetailsExe(req)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "[BlockDetails] query error", "查询错误"))
	}
	log.Debugf("[BlockDetails] query res success,res: %v", resWeb)
	return c.JSON(pkg.SuccessResponse(resWeb))
}

func BlockDetailsExe(req types.BlockDetailsReq) (*types.BlockDetailsResWeb, error) {
	ctx, client := es.GetESInstance()
	resWeb, err := es.BlockDetailsProcessed(ctx, client, int(req.QueryType), req.QueryWord)
	if err != nil {
		log.Errorf("[BlockDetails] Es BlockDetailsProcessed QueryWord: %v, error: %s", req.QueryWord, err)
		return nil, err
	}
	return resWeb, nil
	//res, err := es.GetBlockDetails(req.QueryType, req.QueryWord)
	//if err != nil {
	//	log.Errorf("[BlockDetails] Es Block QueryWord: %v, error: %s", req.QueryWord, err)
	//	return nil, err
	//}
	//if res == nil {
	//	log.Error("[BlockDetails] res nil")
	//	return nil, err
	//}
	//cRes, err := es.QueryChunkDetails(ctx, client, pkg.ChunkQueryType(req.QueryType), req.QueryWord)
	//if err != nil {
	//	log.Errorf("[BlockDetails] QueryChunkDetails KeyWord: %s, Error: %s", req.QueryWord, err)
	//	return nil, err
	//}
	//gu := pkg.DivisionPowerOfTen(float64(res.GasPrice), 9)
	//gl := pkg.DivisionPowerOfTen(float64(res.GasLimit), 15)
	//resWeb := types.BlockDetailsResWeb{
	//	Height:           res.Height,
	//	Hash:             res.Hash,
	//	Timestamp:        res.Timestamp,
	//	TimestampNanoSec: res.TimestampNanoSec,
	//	Transactions:     int64(len(cRes.Transactions)),
	//	Receipts:         int64(len(cRes.Receipts)),
	//	Author:           res.Author,
	//	GasUsed:          res.GasUsed,
	//	GasPrice:         gu,
	//	GasLimit:         gl,
	//	GasFee:           float64(res.GasUsed) * gu,
	//	PrevHash:         res.PrevHash,
	//}
}

// @Tags Web
// @Summary [Block] LastBlock
// @Accept json
// @Description LastBlock API
// @Param param body nil false "Request Params"
// @Success 200 {object} types.LastBlockRes "Success Response"
// @Router /block/last [post]
func LastBlock(c *fiber.Ctx) error {
	res, _ := es.GetLastBlocks()
	log.Debugf("BlockDetails res success, res: %v", res)
	return c.JSON(pkg.SuccessResponse(res))
}

// @Tags Web
// @Summary [Block] BlockChanges
// @Accept json
// @Description BlockChanges API
// @Param param body string false "block_hash"
// @Success 200 {object} types.BlockChangesBody "Success Response"
// @Router /block/changes [post]
func FinalBlockChanges(c *fiber.Ctx) error {
	var req types.BlockHashReq
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	}
	res, _ := es.QueryFinalBlockChanges(req.BlockHash)
	fmt.Println("BlockDetails res success...")
	return c.JSON(pkg.SuccessResponse(res))
}

func BlockList(c *fiber.Ctx) error {
	numStr := c.Query("page_num", "1")
	sizeStr := c.Query("page_size", "20")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), "页码错误"))
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), "页码错误"))
	}
	if num <= 1 {
		num = 1
	}
	if size <= 0 {
		size = 20
	}
	ctx, client := es.GetESInstance()
	esRes, total, err := es.QueryBlockList(ctx, client, num, size)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), "查询错误"))
	}
	res := types.BlockListResWeb{
		Total:     total,
		BlockList: esRes,
	}
	return c.JSON(pkg.SuccessResponse(res))
}
