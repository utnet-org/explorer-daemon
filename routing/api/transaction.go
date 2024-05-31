package api

import (
	"explorer-daemon/es"
	"explorer-daemon/pkg"
	"explorer-daemon/types"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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
		//b, err := es.GetBlockDetails(pkg.BlockQueryHash, v.TransactionOutcome.BlockHash)
		//if err != nil {
		//	return c.JSON(pkg.MessageResponse(-1, err.Error(), "数据不存在"))
		//}
		tb, _ := pkg.DivisionBigPowerOfTen(v.TransactionOutcome.Outcome.TokensBurnt, 24)
		result = append(result, types.TxnResWeb{
			Height:     v.Height,
			Timestamp:  v.Timestamp,
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

func GetTxnDetail(c *fiber.Ctx) error {
	txnHash := gjson.Get(string(c.Body()), "txn_hash").String()
	resWeb, err := GetTxnDetailExe(txnHash)
	if err != nil {
		return c.JSON(pkg.MessageResponse(pkg.MESSAGE_FAIL, "[BlockDetails] query error", "查询错误"))
	}
	log.Debugf("[BlockDetails] query res success,res: %v", resWeb)
	return c.JSON(pkg.SuccessResponse(resWeb))
}

func GetTxnDetailExe(txnHash string) (*types.TxnDetailResWeb, error) {
	ctx, client := es.GetESInstance()
	esRes, err := es.QueryTxnStatusByHash(ctx, client, txnHash)
	if err != nil {
		log.Errorf("[TxnDetailExe] Es QueryTxnStatusByHash TxnHash: %v, error: %s", txnHash, err)
		return nil, err
	}
	tb, _ := pkg.DivisionBigPowerOfTen(esRes.TransactionOutcome.Outcome.TokensBurnt, 24)
	resWeb := &types.TxnDetailResWeb{
		Hash:             esRes.Transaction.Hash,
		Status:           isTransactionSuccessful(esRes.TransactionOutcome.Outcome.Status),
		Height:           esRes.Height,
		Timestamp:        esRes.Timestamp,
		TimeUTC:          pkg.NanoToUTCStr(esRes.Timestamp),
		SignerID:         esRes.Transaction.SignerID,
		ReceiverID:       esRes.Transaction.ReceiverID,
		TokenTransferred: nil,
		Deposit:          "",
		TxnFee:           tb,
	}
	return resWeb, nil
}

func isTransactionSuccessful(txStatus interface{}) string {
	//status := txStatus.TransactionOutcome.Outcome.Status
	if statusMap, ok := txStatus.(map[string]interface{}); ok {
		if _, ok := statusMap["SuccessReceiptId"]; ok {
			//return "SuccessValue"
			return statusMap["SuccessReceiptId"].(string)
		}
		if _, ok := statusMap["Failure"]; ok {
			return "Failure"
		}
	}
	return ""
}

func GetAccountTxns(c *fiber.Ctx) error {
	accId := c.Query("account_id")
	num := c.QueryInt("page_num")
	size := c.QueryInt("page_size")
	if num <= 1 {
		num = 1
	}
	if size <= 0 {
		size = 10
	}
	ctx, client := es.GetESInstance()
	esRes, total, err := es.QueryAccountTxns(ctx, client, num, size, accId)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-1, err.Error(), "查询交易失败"))
	}

	var result []types.TxnResWeb
	for _, v := range esRes {
		tb, _ := pkg.DivisionBigPowerOfTen(v.TransactionOutcome.Outcome.TokensBurnt, 24)
		result = append(result, types.TxnResWeb{
			Height:     v.Height,
			Timestamp:  v.Timestamp,
			Hash:       v.Transaction.Hash,
			TxnType:    "", // unknown
			ReceiverID: v.Transaction.ReceiverID,
			SignerID:   v.Transaction.SignerID,
			Deposit:    "", // only function call has deposit
			TxnFee:     tb,
		})
	}
	webRes := types.TxnListResWeb{
		Total:   total,
		TxnList: result,
	}
	return c.JSON(pkg.SuccessResponse(webRes))
}
