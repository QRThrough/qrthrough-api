package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const LINE_PREFIX = "/line"

func bindLineBotRouter(router fiber.Router) {
	lineRouter := router.Group(LINE_PREFIX)

	qrRepo := repo.NewQRCodeRepo(infrastructure.DB)
	accRepo := repo.NewAccountRepo(infrastructure.DB)
	configurationRepo := repo.NewConfigurationRepo(infrastructure.DB)
	serv := service.NewLineService(qrRepo, accRepo, configurationRepo)
	hdl := handler.NewLineHandler(serv)

	lineRouter.Post("/webhook", hdl.Webhook)
	// lineRouter.Post("/qrcode", hdl.ManualQR)
}
