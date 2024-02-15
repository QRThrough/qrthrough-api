package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/JMjirapat/qrthrough-api/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

const DASHBOARD_PREFIX = "/dashboard"

func bindDashboardRouter(router fiber.Router) {
	dashboard := router.Group(DASHBOARD_PREFIX)

	accountRepo := repo.NewAccountRepo(infrastructure.DB)
	configurationRepo := repo.NewConfigurationRepo(infrastructure.DB)
	serv := service.NewDashboardService(accountRepo, configurationRepo)
	hdl := handler.NewDashboardHandler(serv)

	roleLimit := []model.Role{model.ROLE_ADMIN, model.ROLE_MODERATOR}
	dashboard.Use(middlewares.AuthMiddleware(serv))
	dashboard.Use(middlewares.AccessControlMiddleware(roleLimit))
	dashboard.Get("/signin", hdl.SignIn)
	dashboard.Get("/configurations", hdl.AllConfiguration)
	dashboard.Put("/configurations", hdl.UpdateConfiguration)
}
