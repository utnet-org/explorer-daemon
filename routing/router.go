package routing

import (
	"explorer-daemon/routing/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(f *fiber.App) {
	prep := f.Group("/api")

	//overview
	prep.Post("/overview", api.OverviewInfo)

	//search
	prep.Post("/query/filter", api.QueryFilter)

	//block
	prep.Post("/block/last", api.LastBlock)
	prep.Post("/block/details", api.BlockDetails)
	prep.Post("/block/changes", api.FinalBlockChanges)
	prep.Get("/block/list", api.BlockList)

	//chart
	prep.Post("/power/chart", api.PowerChart)

	//gas
	prep.Post("/gas/chart", api.GasChart)

	//network
	prep.Post("/network/validators", api.GetValidators)
	prep.Post("/network/validator", api.GetValidator)

	//chip
	prep.Post("add/chipinfo", api.AddChipInfo)
	prep.Post("query/chipinfo", api.QueryChipInfo)
	prep.Get("chip/list", api.GetChipList)

	//transaction
	prep.Get("txn/list", api.GetTxnList)
	prep.Post("txn/detail", api.GetTxnDetail)

	//contract
	prep.Post("contract/detail", api.ContractDetail)

	// coin
	prep.Get("/coin/price", api.CoinPrice)
}
