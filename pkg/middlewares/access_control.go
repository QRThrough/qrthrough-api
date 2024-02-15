package middlewares

import (
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

func AccessControlMiddleware(roleLimit []model.Role) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		forbidden := errors.NewForbiddenError("Forbidden")
		access := false

		roleStr := c.Locals("ROLE").(string)
		if roleStr == "" {
			return rest.ResponseError(c, forbidden)
		}

		role := model.Role(roleStr)
		for _, val := range roleLimit {
			if val == role {
				access = true
			}
		}

		if !access {
			return rest.ResponseError(c, forbidden)
		}

		return c.Next()
	}
}
