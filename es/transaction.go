package es

import (
	"context"
	"encoding/json"
	"explorer-daemon/types"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

func InsertTxnStatus(ctx context.Context, client *elastic.Client, result types.TxnStatusResult) error {
	_, err := client.Index().
		Index("transaction").
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

func QueryTxnStatusList(ctx context.Context, client *elastic.Client, pageNum, pageSize int) ([]types.TxnStatusResult, int64, error) {
	from := (pageNum - 1) * pageSize
	searchResult, err := client.Search().
		Index("transaction").
		From(from).
		Size(pageSize).
		Do(ctx)
	if err != nil {
		return nil, 0, err
	}

	var txns []types.TxnStatusResult
	for _, hit := range searchResult.Hits.Hits {
		var res types.TxnStatusResult
		if err := json.Unmarshal(hit.Source, &res); err == nil {
			txns = append(txns, res)
		}
	}
	return txns, searchResult.TotalHits(), nil
}
