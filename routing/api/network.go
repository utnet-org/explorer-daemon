package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func GetValidator(c *fiber.Ctx) error {
	accId := gjson.Get(string(c.Body()), "account_id").String()
	ctx, client := es.GetESInstance()
	vs, err := es.QueryValidator(ctx, client)
	if err != nil {
		log.Errorf("[GetValidator] Es Error: %s", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), ""))
	}
	for _, v := range vs.CurrentValidators {
		if accId == v.AccountID {
			return c.JSON(pkg.SuccessResponse(v))
		}
	}
	return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "validator not exist", "账户不存在"))
}

func GetValidators(c *fiber.Ctx) error {
	ctx, client := es.GetESInstance()
	vs, err := es.QueryValidator(ctx, client)
	if err != nil {
		log.Errorf("[GetValidators] Es Error: %s", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, err.Error(), ""))
	}
	return c.JSON(pkg.SuccessResponse(vs))
}
