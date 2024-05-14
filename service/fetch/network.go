package fetch

import (
	"explorer-daemon/es"
	"explorer-daemon/service/remote"
	"fmt"
)

func HandleNetworkInfo() error {
	res, err := remote.NetworkInfo()
	if err != nil {
		fmt.Println("rpc error")
		return err
	}
	if res == nil {
		fmt.Println("rpc error res nil")
		return err
	}
	ctx, client := es.GetESInstance()
	err = es.InsertNetworkInfo(ctx, client, res.Result)
	if err != nil {
		fmt.Println("[HandleNetworkInfo] insert error:", err)
		return err
	}
	return nil
}

func HandleValidation() error {
	res, err := remote.ValidationStatusByNull()
	if err != nil {
		fmt.Println("rpc error")
		return err
	}
	if res == nil {
		fmt.Println("rpc error res nil")
		return err
	}
	//ctx, client := es.GetESInstance()
	//err = es.InsertNetworkInfo(ctx, client, res.Result)
	//if err != nil {
	//	fmt.Println("[HandleNetworkInfo] insert error:", err)
	//	return err
	//}
	return nil
}
