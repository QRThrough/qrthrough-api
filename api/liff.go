package api

import (
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/handler"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const LIFF_PREFIX = "/liff"

func bindLiffRouter(router fiber.Router) {
	liffRouter := router.Group(LIFF_PREFIX)

	accRepo := repo.NewAccountRepo(infrastructure.DB)
	alumniRepo := repo.NewAlumniRepo(infrastructure.DB)
	serv := service.NewLiffService(accRepo, alumniRepo)
	hdl := handler.NewLiffHandler(serv)

	liffRouter.Get("/alumni/:id", hdl.GetAlumni)
	liffRouter.Post("/user", hdl.SignUp)
	liffRouter.Post("/otp/request", hdl.GetOTP)
	liffRouter.Put("/otp/verify", hdl.VerifyOTP)
}
