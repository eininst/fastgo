package v1

import (
	"github.com/gofiber/fiber/v2"
)

type request struct {
	RequestField string
}

type response struct {
	ResponseField string
}

// @Param request body v1.request true "query params"
// @Success 200 {object} v1.response
// @Router /test [post]
func HelloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"name": "hello wolrd!"})
}
