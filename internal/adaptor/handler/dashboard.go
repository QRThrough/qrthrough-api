package handler

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

type dashboardHandler struct {
	serv domain.DashboardService
}

func NewDashboardHandler(serv domain.DashboardService) *dashboardHandler {
	return &dashboardHandler{
		serv: serv,
	}
}

func (h dashboardHandler) SignIn(c *fiber.Ctx) error {
	forbidden := errors.NewForbiddenError("Forbidden")
	roleStr := c.Locals("ROLE").(string)
	if roleStr == "" {
		return rest.ResponseError(c, forbidden)
	}

	result := dto.SignInResponseBody{
		Role: model.Role(roleStr),
	}

	return rest.ResponseOK(c, result)
}

func (h dashboardHandler) AllConfiguration(c *fiber.Ctx) error {
	result, err := h.serv.AllConfiguration()
	if err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}
	return rest.ResponseOK(c, result)
}

func (h dashboardHandler) UpdateConfiguration(c *fiber.Ctx) error {
	var body []dto.Configuration
	if err := c.BodyParser(&body); err != nil {
		return rest.ResponseUnprocessableEntity(c)
	}

	if err := h.serv.UpdateConfiguration(body); err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}
	return rest.ResponseOK(c, nil)
}
