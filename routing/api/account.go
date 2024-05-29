package api

import (
	"errors"
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func AccountDetail(c *fiber.Ctx) error {
	accId := gjson.Get(string(c.Body()), "account_id").String()
	res, err := AccountDetailExe(accId)
	if err != nil {
		log.Errorf("[ViewAccount] AccountDetailExe error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	return c.JSON(pkg.SuccessResponse(res))
}

func AccountDetailExe(accId string) (*types.AccountResult, error) {
	res, err := remote.ViewAccount(accId)
	if err != nil {
		log.Errorf("[AccountDetailExe] ViewAccount error: %v", err)
		return nil, err
	}
	log.Debugf("AccountDetailExe res success, res: %v", res)
	return nil, errors.New("account not exist")
}

func ContractDetail(c *fiber.Ctx) error {
	res, err := remote.ViewContractCode("unc")
	if err != nil {
		log.Errorf("[ContractDetail] ViewContractCode error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	log.Debugf("ContractDetail res success, res: %v", res)
	return c.JSON(pkg.SuccessResponse(res.CodeBase64))
}
