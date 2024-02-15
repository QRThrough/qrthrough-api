package handler

import (
	"strconv"

	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

type scannerHandler struct {
	serv domain.ScannerService
}

func NewScannerHandler(serv domain.ScannerService) *scannerHandler {
	return &scannerHandler{
		serv: serv,
	}
}

func (h scannerHandler) VerifyQR(c *fiber.Ctx) error {
	forbidden := errors.NewForbiddenError("Forbidden")

	if !h.serv.CheckExistedScanner(c.Get("Authorization")) {
		return rest.ResponseError(c, forbidden)
	}

	qrCodeId := c.Params("token")
	id, err := strconv.ParseInt(qrCodeId, 10, 64)
	if err != nil {
		return rest.ResponseBadRequest(c)
	}

	if err = h.serv.VerifyQR(id); err != nil {
		return rest.ResponseError(c, errors.NewUnauthorizedError(err.Error()))
	}

	return rest.ResponseOK(c, nil)
}
