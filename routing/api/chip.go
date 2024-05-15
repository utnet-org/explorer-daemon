package api

import (
	"crypto/sha256"
	"encoding/hex"
	"explorer-daemon/database"
	"explorer-daemon/model"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Tags Web
// @Summary [Overview] AddChipInfo
// @Accept json
// @Description AddChipInfo API
// @Param param body types.AddChipInfoReq true "Request Params"
// @Success 200 {object} types.AddChipInfoRes "Success Response"
// @Router /add/chipinfo [post]

func AddChipInfo(c *fiber.Ctx) error {
	reqParams := types.AddChipInfoReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return err
	}
	chip, err := model.GetChipBySerialBus(database.DB, reqParams.SerialNumber, reqParams.BusId)
	if chip != nil {
		return c.JSON(pkg.MessageResponse(-1, "repeated chip", "芯片已经存在！"))
	}

	// SearchKey 生成
	seed := reqParams.SerialNumber + reqParams.PublicKey[:5]
	hash := sha256.New()
	hash.Write([]byte(seed))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)[:20]

	newUser := model.Chip{
		Model:        gorm.Model{},
		SearchKey:    "UTC" + hashString,
		ChipType:     reqParams.ChipType,
		Power:        reqParams.Power,
		SerialNumber: reqParams.SerialNumber,
		BusId:        reqParams.BusId,
		P2Key:        reqParams.P2Key,
		PubKey:       reqParams.PublicKey,
		Flag:         "1",
	}
	_, err = newUser.InsertNewChip(database.DB)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-1, err.Error(), "插入新芯片失败"))
	}

	return c.JSON(pkg.SuccessResponse("add chip information success"))
}

// @Tags Web
// @Summary [Overview] QueryChipInfo
// @Accept json
// @Description QueryChipInfo API
// @Param param body types.QueryChipInfoReq true "Request Params"
// @Success 200 {object} types.QueryChipInfoRes "Success Response"
// @Router /query/chipinfo [post]

func QueryChipInfo(c *fiber.Ctx) error {
	reqParams := types.QueryChipInfoReq{}
	err := c.BodyParser(&reqParams)
	if err != nil {
		return err
	}
	result, err := QueryChipInfoExe(reqParams.SearchKey)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-1, err.Error(), "查询芯片失败"))
	}
	return c.JSON(pkg.SuccessResponse(result))
}

func QueryChipInfoExe(keyword string) ([]types.QueryChipInfoRep, error) {
	chips, err := model.GetChipBySearchKey(database.DB, keyword)
	if err != nil {
		return nil, err
	}
	result := make([]types.QueryChipInfoRep, 0)
	for _, item := range chips {
		result = append(result, types.QueryChipInfoRep{
			ChipType:     item.ChipType,
			Power:        item.Power,
			SerialNumber: item.SerialNumber,
			BusId:        item.BusId,
			P2Key:        item.P2Key,
			PublicKey:    item.PubKey,
		})
	}
	return result, nil
}
