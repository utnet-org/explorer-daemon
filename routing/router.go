package routing

import (
	"explorer-daemon/routing/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(f *fiber.App) {
	prep := f.Group("/api")

	//api.Post("/example", app.Example)
	prep.Post("/overview", api.OverviewInfo)

	prep.Post("/query/filter", api.QueryFilter)

	// block
	prep.Post("/block/last", api.LastBlock)
	prep.Post("/block/details", api.BlockDetails)
	prep.Post("/block/changes", api.FinalBlockChanges)

	prep.Post("/power/chart", api.PowerChart)

	prep.Post("/gas/chart", api.GasChart)

	prep.Post("/network/validators", api.GetValidators)
	prep.Post("/network/validator", api.GetValidator)
}
