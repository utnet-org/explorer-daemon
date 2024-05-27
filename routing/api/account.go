package api

import (
	"explorer-daemon/pkg"
	"explorer-daemon/service/remote"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func ContractDetail(c *fiber.Ctx) error {
	res, err := remote.ViewContractCode("unc")
	if err != nil {
		log.Errorf("[ContractDetail] ViewContractCode error: %v", err)
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "error", ""))
	}
	log.Debugf("ContractDetail res success, res: %v", res)
	return c.JSON(pkg.SuccessResponse(res.CodeBase64))
}
