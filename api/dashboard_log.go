package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const DASHBOARD_LOG_PREFIX = DASHBOARD_PREFIX + "/logs"

func bindDashboardLogRouter(router fiber.Router) {
	logRouter := router.Group(DASHBOARD_LOG_PREFIX)

	logRepo := repo.NewLogRepo(infrastructure.DB)
	serv := service.NewDashboardLogService(logRepo)
	hdl := handler.NewDashboardLogHandler(serv)

	logRouter.Get("", hdl.All)
}
