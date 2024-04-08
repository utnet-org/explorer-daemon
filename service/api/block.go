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
	res := es.LastBlockQuery()
	resBody := types.LastBlockResList{
		LastBlockList: res,
	}
	fmt.Println("BlockDetails res success...")
	pkg.PrintStruct(resBody)
	return c.JSON(pkg.SuccessResponse(resBody))
}
