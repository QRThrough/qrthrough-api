package api

import (
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/gofiber/fiber/v2"
)

func bindHealthRouter(router fiber.Router) {
	router.Get("", func(c *fiber.Ctx) error {
		return rest.ResponseOK(c, nil)
	})
}
