package es

import (
	"context"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertGas(ctx context.Context, client *elastic.Client, result types.GasStoreResult) error {
	_, err := client.Index().
		Index("gas").
		Id(result.Hash).
		BodyJson(result).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Debugln("[InsertTransaction] Success")
	return nil
}
