package handler

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

type dashboardModeratorHandler struct {
	serv domain.DashboardModeratorService
}

func NewDashboardModeratorHandler(serv domain.DashboardModeratorService) *dashboardModeratorHandler {
	return &dashboardModeratorHandler{
		serv: serv,
	}
}

func (h dashboardModeratorHandler) All(c *fiber.Ctx) error {
	query := domain.DashboardUserQuery{}
	if err := c.QueryParser(&query); err != nil {
		return rest.ResponseBadRequest(c)
	}

	result, err := h.serv.All(query)
	if err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseOK(c, result)
}

func (h dashboardModeratorHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return rest.ResponseBadRequest(c)
	}

	var body dto.AccountRequestBody
	if err := c.BodyParser(&body); err != nil {
		return rest.ResponseUnprocessableEntity(c)
	}

	if err := h.serv.Update(id, body); err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseOK(c, nil)
}

func (h dashboardModeratorHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return rest.ResponseBadRequest(c)
	}

	if err := h.serv.Delete(id); err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseOK(c, nil)
}
