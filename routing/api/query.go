package api

import (
	"explorer-daemon/pkg"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type QueryConditionReq struct {
	Condition string         `json:"condition"`
	QueryType pkg.SearchType `json:"query_type"`
}

// ExecuteQuery 方法
func (q QueryConditionReq) ExecuteQuery() string {
	switch q.QueryType {
	case pkg.SearchAccount:
		fmt.Printf("account")
	case pkg.SearchBlock:
		fmt.Printf("block")

	default:
		fmt.Printf("default")
	}
	return fmt.Sprintf("Executing account query: %s", q.Condition)
}

// @Tags Web
// @Summary [Query] QueryFilter
// @Accept json
// @Description QueryFilter API
// @Param param body types.Example false "Request Params"
// @Success 200 {object} types.ExampleRes "Success Response"
// @Router /query/filter [post]
func QueryFilter(c *fiber.Ctx) error {
	var req QueryConditionReq
	err := c.BodyParser(&req)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "can not transfer request to struct", "请求参数错误"))
	}
	// 根据不同类型的查询条件执行相应的查询操作
	queryResult := req.ExecuteQuery()
	return c.JSON(pkg.SuccessResponse(queryResult))
}
