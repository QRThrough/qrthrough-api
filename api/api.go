package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitAPI(app *fiber.App) {
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(logger.New())
	app.Use(recover.New())

	path := app.Group("/api")
	bindHealthRouter(path)
	bindDeviceRouter(path)
	bindLineBotRouter(path)
	bindLiffRouter(path)
	bindDashboardRouter(path)
	bindDashboardUserRouter(path)
	bindDashboardLogRouter(path)
	bindDashboardModeratorRouter(path)
}
