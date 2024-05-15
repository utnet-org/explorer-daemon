package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
)

func HandleChipQuery() error {
	res, err := remote.ChipsQuery()
	if err != nil {
		log.Errorf("[HandleChipQuery] rpc error: %v", err)
		return err
	}
	ctx, client := es.GetESInstance()
	qRes, err := es.QueryChipByHeight(ctx, client, res.Result.BlockHeight)
	if err != nil {
		log.Errorf("[HandleChipQuery] query error: %v", err)
		return err
	}
	if qRes.TotalHits() > 0 {
		log.Infof("[HandleChipQuery] Chip data exist, height: %v", res.Result.BlockHeight)
		return nil
	}
	err = es.InsertChip(ctx, client, res.Result)
	if err != nil {
		log.Errorf("[HandleChipQuery] insert error: %v", err)
		return err
	}
	return nil
}
