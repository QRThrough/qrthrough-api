package handler

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

type dashboardLogHandler struct {
	serv domain.DashboardLogService
}

func NewDashboardLogHandler(serv domain.DashboardLogService) *dashboardLogHandler {
	return &dashboardLogHandler{
		serv: serv,
	}
}

func (h dashboardLogHandler) All(c *fiber.Ctx) error {
	query := domain.DashboardLogQuery{}
	if err := c.QueryParser(&query); err != nil {
		return rest.ResponseBadRequest(c)
	}

	result, err := h.serv.All(query)
	if err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseOK(c, result)
}
