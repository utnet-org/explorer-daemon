package es

import (
	"context"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertAccount(ctx context.Context, client *elastic.Client, result types.TxnStoreResult) error {
	_, err := client.Index().
		Index("account").
		Id(result.Transaction.Hash).
		BodyJson(result).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Debugln("[InsertTransaction] Success")
	return nil
}
