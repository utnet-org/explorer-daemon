package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
)

func GetTxnList(c *fiber.Ctx) error {
	num := c.QueryInt("page_num", 1)
	size := c.QueryInt("page_size", 20)
	if num <= 1 {
		num = 1
	}
	if size <= 0 {
		size = 20
	}
	ctx, client := es.GetESInstance()
	esRes, total, err := es.QueryTxnStatusList(ctx, client, num, size)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-1, err.Error(), "查询交易失败"))
	}

	var result []types.TxnResWeb
	for _, v := range esRes {
		b, err := es.GetBlockDetails(pkg.BlockQueryHash, v.TransactionOutcome.BlockHash)
		if err != nil {
			return c.JSON(pkg.MessageResponse(-1, err.Error(), "数据不存在"))
		}
		tb, _ := pkg.DivisionBigPowerOfTen(v.TransactionOutcome.Outcome.TokensBurnt, 24)
		result = append(result, types.TxnResWeb{
			Height:     v.Height,
			Timestamp:  b.Timestamp,
			Hash:       v.Transaction.Hash,
			TxnType:    "", // unknown
			ReceiverID: v.Transaction.ReceiverID,
			SignerID:   v.Transaction.SignerID,
			Deposit:    "", // only function call has deposit
			//TxnFee:     strconv.FormatFloat(tb, 'f', -1, 64),
			TxnFee: tb,
		})
	}
	webRes := types.TxnListResWeb{
		Total:   total,
		TxnList: result,
	}
	return c.JSON(pkg.SuccessResponse(webRes))
}
