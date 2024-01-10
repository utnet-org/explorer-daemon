package routing

import (
	"explorer-daemon/service/app"
	"explorer-daemon/service/web"
	"github.com/gofiber/fiber/v2"
)

func Setup(f *fiber.App) {
	api := f.Group("/api")

	api.Post("/example", app.Example)
	api.Post("/overview/info", web.OverviewInfo)
}
