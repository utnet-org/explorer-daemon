package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	log "github.com/sirupsen/logrus"
)

func HandleMiner() error {
	ctx, client := es.GetESInstance()
	last, err := es.GetLastHeightHash(client, ctx)
	if err != nil {
		log.Errorf("[HandleMiner] GetLastHeight error: %v\n", err)
		return err
	}
	res, err := remote.AllMiners(last.Hash)
	if err != nil {
		return err
	}
	if err = es.InsertMiner(ctx, client, res.Result); err != nil {
		log.Error("[HandleMiner] InsertBlockDetails error:", err)
		return err
	}
	return nil
}
