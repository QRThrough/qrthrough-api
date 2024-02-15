package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const DASHBOARD_USER_PREFIX = DASHBOARD_PREFIX + "/users"

func bindDashboardUserRouter(router fiber.Router) {
	userRouter := router.Group(DASHBOARD_USER_PREFIX)

	accountRepo := repo.NewAccountRepo(infrastructure.DB)
	serv := service.NewDashboardUserService(accountRepo)
	hdl := handler.NewDashboardUserHandler(serv)

	userRouter.Get("", hdl.All)
	userRouter.Put("/:id", hdl.Update)
	userRouter.Delete("/:id", hdl.Delete)
}
