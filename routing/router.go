package routing

import (
	"explorer-daemon/service/web"
	"github.com/gofiber/fiber/v2"
)

func Setup(f *fiber.App) {
	api := f.Group("/api")

	//api.Post("/example", app.Example)
	api.Post("/overview/info", web.OverviewInfo)

	// block
	api.Post("/block/details", web.BlockDetails)
}
