package es

import (
	"explorer-daemon/types"
	"fmt"
)

func InsertChip(result types.ChipQueryResult) error {
	//var cs = make([]types.Chip, 0)
	//cs = append(cs, types.Chip{
	//	MinerId:   "1",
	//	Power:     1,
	//	BusId:     "123",
	//	PublicKey: "456",
	//	ChipSN:    "789",
	//	P2Key:     "123",
	//})
	//result.Chip = cs
	createIndexIfNotExists(ECTX, ECLIENT, "chip")
	_, err := ECLIENT.Index().
		Index("chip").
		BodyJson(&result).
		Do(ECTX)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("InsertChip Success")
	return nil
}
