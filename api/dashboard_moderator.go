package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const DASHBOARD_MODERATOR_PREFIX = DASHBOARD_PREFIX + "/moderators"

func bindDashboardModeratorRouter(router fiber.Router) {
	moderatorRouter := router.Group(DASHBOARD_MODERATOR_PREFIX)

	accountRepo := repo.NewAccountRepo(infrastructure.DB)
	serv := service.NewDashboardModeratorService(accountRepo)
	hdl := handler.NewDashboardModeratorHandler(serv)

	moderatorRouter.Get("", hdl.All)
	moderatorRouter.Put("/:id", hdl.Update)
	moderatorRouter.Delete("/:id", hdl.Delete)
}
