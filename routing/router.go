package routing

import (
	"explorer-daemon/service/api"
	"github.com/gofiber/fiber/v2"
)

func Setup(f *fiber.App) {
	prep := f.Group("/api")

	//api.Post("/example", app.Example)
	prep.Post("/overview/info", api.OverviewInfo)

	// block
	prep.Post("/block/last", api.LastBlock)
	prep.Post("/block/details", api.BlockDetails)
}
