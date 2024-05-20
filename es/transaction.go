package es

import (
	"context"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertTransaction(ctx context.Context, client *elastic.Client, result types.Transaction) error {
	_, err := client.Index().
		Index("transaction").
		Id("network").
		BodyJson(result).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Debugln("InsertNetWorkInfo Success")
	return nil
}
