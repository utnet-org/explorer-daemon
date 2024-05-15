package api

import (
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

type QueryConditionReq struct {
	Keyword   interface{}    `json:"keyword"`
	QueryType pkg.SearchType `json:"query_type"`
}

type QueryConditionRes struct {
	Data      interface{}
	QueryType pkg.SearchType `json:"query_type"`
}

func (q QueryConditionReq) ExecuteQuery() interface{} {
	switch q.QueryType {
	case pkg.SearchBlockHeight:
		req := types.BlockDetailsReq{
			QueryType: pkg.BlockQueryHeight,
			QueryWord: q.Keyword,
		}
		res, err := BlockDetailsExe(req)
		if err != nil {
			return nil
		}
		log.Infof("[ExecuteQuery] BlockDetailsExe:%v", res)
		return res
	case pkg.SearchBlockHash:
		req := types.BlockDetailsReq{
			QueryType: pkg.BlockQueryHash,
			QueryWord: q.Keyword,
		}
		res, err := BlockDetailsExe(req)
		if err != nil {
			return nil
		}
		log.Infof("[ExecuteQuery] BlockDetailsExe:%v", res)
		return res
	case pkg.SearchAccount:
		fmt.Printf("account")
	case pkg.SearchChip:
		fmt.Printf("chip")

	default:
		fmt.Printf("default")
	}
	return fmt.Sprintf("Executing account query: %s", q.Keyword)
}

func ExecuteQuery2(keyword interface{}) (interface{}, error) {
	numberRegex := regexp.MustCompile(`^\d+$`)
	strKey := fmt.Sprintf("%v", keyword)
	if numberRegex.MatchString(strKey) {
		intKey, err := strconv.ParseInt(strKey, 10, 64)
		if err != nil {
			return nil, err
		}
		req := types.BlockDetailsReq{
			QueryType: pkg.BlockQueryHeight,
			QueryWord: intKey,
		}
		res, err := BlockDetailsExe(req)
		if err != nil {
			return nil, err
		}
		log.Debugf("[ExecuteQuery] BlockDetailsExe:%v", res)
		return pkg.QueryResponse(res, pkg.SearchBlockHeight), nil
	}
	if len(strKey) == 44 {
		req := types.BlockDetailsReq{
			QueryType: pkg.BlockQueryHash,
			QueryWord: keyword,
		}
		res, err := BlockDetailsExe(req)
		if err != nil {
			return nil, err
		}
		log.Debugf("[ExecuteQuery] BlockDetailsExe:%v", res)
		return pkg.QueryResponse(res, pkg.SearchBlockHash), nil
	}
	if len(strKey) == 64 {
		res, err := GetValidatorExe(strKey)
		if err != nil {
			return nil, err
		}
		log.Debugf("[ExecuteQuery] GetValidatorExe:%v", res)
		return pkg.QueryResponse(res, pkg.SearchAccount), nil
	}
	if strings.HasPrefix(strings.TrimSpace(strKey), "UTC") {
		res, err := QueryChipInfoExe(strKey)
		if err != nil {
			return nil, err
		}
		log.Debugf("[ExecuteQuery] QueryChipInfoExe:%v", res)
		return pkg.QueryResponse(res, pkg.SearchChip), nil
	}
	return nil, nil
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
	//queryResult := req.ExecuteQuery()
	queryResult, err := ExecuteQuery2(req.Keyword)
	if err != nil {
		log.Errorf("[QueryFilter] Error:%v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), "查询出错"))
	}
	return c.JSON(queryResult)
}
