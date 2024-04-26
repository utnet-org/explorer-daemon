package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
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
	resBody := &types.BlockDetailsBody{}

	resBody, err = es.BlockDetailsQuery(req.QueryWord, pkg.BlockQueryType(req.QueryType))
	fmt.Println("[BlockDetails] query res success")
	pkg.PrintStruct(resBody)
	return c.JSON(pkg.SuccessResponse(resBody))
}

// @Tags Web
// @Summary [Block] LastBlock
// @Accept json
// @Description LastBlock API
// @Param param body nil false "Request Params"
// @Success 200 {object} types.LastBlockRes "Success Response"
// @Router /block/last [post]
func LastBlock(c *fiber.Ctx) error {
	res, _ := es.LastBlockQuery()
	//var blocks []types.LastBlockRes
	//for index, hit := range res {
	//	var body types.LastBlockRes
	//	fmt.Printf("第 %d 条数据\n", index+1)
	//	_ = json.Unmarshal(hit.Source, &body)
	//	pkg.PrintStruct(body)
	//	body.Msgs = 1
	//	blocks = append(blocks, body)
	//}
	resBody := types.LastBlockResList{
		LastBlockList: *res,
	}
	fmt.Println("BlockDetails res success...")
	pkg.PrintStruct(resBody)
	return c.JSON(pkg.SuccessResponse(resBody))
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
