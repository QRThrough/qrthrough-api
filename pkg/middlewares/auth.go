package middlewares

import (
	"strings"

	URL "net/url"

	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/internal/core/dto"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(serv domain.DashboardService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// For Development ---->>
		c.Locals("UID", "11234")
		c.Locals("ROLE", "ADMIN")
		return c.Next()
		// For Development <<----
		cfg := config.Config
		invalidToken := errors.NewUnauthorizedError("InvalidToken")
		raw := c.Get("Authorization")

		if raw == "" {
			return rest.ResponseError(c, invalidToken)
		}

		idToken := extractBearerToken(raw)
		if idToken == "" {
			return rest.ResponseError(c, invalidToken)
		}

		url := "https://api.line.me/oauth2/v2.1/verify"
		headers := map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/x-www-form-urlencoded",
		}
		data := URL.Values{}
		data.Add("id_token", idToken)
		data.Add("client_id", cfg.DashboardChannelID)
		res, _, err := rest.HttpPost[dto.Token](url, headers, data.Encode())
		if err != nil {
			return rest.ResponseError(c, errors.NewStatusBadGatewayError("เกิดข้อผิดพลาดกับการติดต่อบริการของ Line"))
		}
		if len(res.Sub) <= 0 {
			return rest.ResponseError(c, errors.NewStatusBadGatewayError("เกิดข้อผิดพลาดกับการติดต่อบริการของ Line"))
		}

		role, err := serv.GetRole(res.Sub)
		if err != nil {
			return rest.ResponseError(c, errors.NewUnauthorizedError("CannotGetRole"))
		}

		c.Locals("UID", res.Sub)
		c.Locals("ROLE", string(role))
		return c.Next()
	}
}

func extractBearerToken(raw string) string {
	return strings.TrimSpace(strings.Replace(raw, "Bearer", "", 1))
}
