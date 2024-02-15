package handler

import (
	"log"
	URL "net/url"

	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

type liffHandler struct {
	serv domain.LiffService
}

func NewLiffHandler(serv domain.LiffService) *liffHandler {
	return &liffHandler{
		serv: serv,
	}
}

func (h liffHandler) GetAlumni(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return rest.ResponseBadRequest(c)
	}

	result, err := h.serv.GetAlumni(id)
	if err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseOK(c, &result)
}

func (h liffHandler) SignUp(c *fiber.Ctx) error {
	cfg := config.Config

	var body dto.RegisterRequestBody
	if err := c.BodyParser(&body); err != nil {
		return rest.ResponseUnprocessableEntity(c)
	}

	url := "https://api.line.me/oauth2/v2.1/verify"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := URL.Values{}
	data.Add("id_token", body.IDToken)
	data.Add("client_id", cfg.ChannelID)
	res, _, err := rest.HttpPost[dto.Token](url, headers, data.Encode())
	if err != nil {
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.LIFF_BAD_GATEWAY_CODE))
	}
	if len(res.Sub) <= 0 {
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.LIFF_NO_SUB_CODE))
	}

	if err := h.serv.SignUp(body, res.Sub); err != nil {
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return rest.ResponseCreated(c, nil)
}

func (h liffHandler) GetOTP(c *fiber.Ctx) error {
	cfg := config.Config

	var body dto.GetOTPRequestBody
	if err := c.BodyParser(&body); err != nil {
		return rest.ResponseUnprocessableEntity(c)
	}

	url := "https://otp.thaibulksms.com/v2/otp/request"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := URL.Values{}
	data.Add("key", cfg.TBSKey)
	data.Add("secret", cfg.TBSSecret)
	data.Add("msisdn", body.Tel)
	res, statusCode, err := rest.HttpPost[dto.OTPReq](url, headers, data.Encode())
	if err != nil || statusCode != 200 {
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.TBS_BAD_GATEWAY_CODE))
	}

	if res.Code != 0 {
		log.Printf("%v : %v", res.Code, res.Errors[0].Message)
		if res.Errors[0].Detail == "ERROR_MSISDN" {
			return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.TBS_MSISDN_ERROR_CODE))
		}
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.TBS_BAD_GATEWAY_CODE))
	}

	result := dto.GetOTPResponseBody{
		Token: res.Token,
		Refno: res.Refno,
	}

	return rest.ResponseCreated(c, result)
}

func (h liffHandler) VerifyOTP(c *fiber.Ctx) error {
	cfg := config.Config

	var body dto.VerifyOTPRequestBody
	if err := c.BodyParser(&body); err != nil {
		return rest.ResponseUnprocessableEntity(c)
	}

	url := "https://otp.thaibulksms.com/v2/otp/verify"
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := URL.Values{}
	data.Add("key", cfg.TBSKey)
	data.Add("secret", cfg.TBSSecret)
	data.Add("token", body.Token)
	data.Add("pin", body.Pin)
	res, _, err := rest.HttpPost[dto.OTPVerify](url, headers, data.Encode())
	if err != nil {
		log.Panicf("%v", err)
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.TBS_BAD_GATEWAY_CODE))
	}

	if res.Code != 0 {
		log.Printf("%v : %v", res.Code, res.Errors[0].Message)
		if res.Errors[0].Message == "Code is invalid." {
			return rest.ResponseError(c, errors.NewUnauthorizedError(errors.TBS_WRONG_OTP_CODE))
		}
		if res.Errors[0].Message == "Token is expire." {
			return rest.ResponseError(c, errors.NewUnauthorizedError(errors.TBS_EXPIRED_OTP_CODE))
		}
		return rest.ResponseError(c, errors.NewStatusBadGatewayError(errors.TBS_BAD_GATEWAY_CODE))
	}

	result := &res.Status

	return rest.ResponseOK(c, result)
}
