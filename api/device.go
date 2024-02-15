package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const DEVICE_PREFIX = "/device"

func bindDeviceRouter(router fiber.Router) {
	deviceRouter := router.Group(DEVICE_PREFIX)

	qrRepo := repo.NewQRCodeRepo(infrastructure.DB)
	logRepo := repo.NewLogRepo(infrastructure.DB)
	scanenrRepo := repo.NewScannerRepo(infrastructure.DB)
	configurationRepo := repo.NewConfigurationRepo(infrastructure.DB)
	serv := service.NewScannerService(qrRepo, logRepo, scanenrRepo, configurationRepo)
	hdl := handler.NewScannerHandler(serv)

	deviceRouter.Get("/scanner/:token", hdl.VerifyQR)
}
