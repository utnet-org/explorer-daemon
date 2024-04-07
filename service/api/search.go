package api

import (
	"explorer-daemon/pkg"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// 定义一个通用的查询条件接口
type QueryCondition interface {
	// 通用的方法，用于执行查询操作
	ExecuteQuery() string
}

type AccountQuery struct {
	Account string `json:"account"`
}

type BlockQuery struct {
	Block string `json:"block"`
}

type AddressQuery struct {
	Address string `json:"address"`
}

// 实现 AccountQuery 结构体的 ExecuteQuery 方法
func (q AccountQuery) ExecuteQuery() string {
	return fmt.Sprintf("Executing account query: %s", q.Account)
}

// 实现 BlockQuery 结构体的 ExecuteQuery 方法
func (q BlockQuery) ExecuteQuery() string {
	return fmt.Sprintf("Executing block query: %d", q.Block)
}

// @Tags Search
// @Summary SearchFilter
// @Accept json
// @Description SearchFilter API
// @Param param body types.Example false "Request Params"
// @Success 200 {object} types.ExampleRes "Success Response"
// @Router /search/filter [post]
func SearchFilter(c *fiber.Ctx) error {
	var req QueryCondition
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	}
	// 根据不同类型的查询条件执行相应的查询操作
	queryResult := req.ExecuteQuery()
	return c.JSON(pkg.SuccessResponse(queryResult))
}
