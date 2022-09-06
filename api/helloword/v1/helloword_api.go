package v1

import (
	"fastgo/internal/service/user"
	"fastgo/pkg/di"
	"github.com/gofiber/fiber/v2"
)

func init() {
	di.Inject(new(HellowordApi))
}

type HellowordApi struct {
	UserService user.UserService `inject:""`
}

// @Summary 测试swagger
// @Tags test
// @version 1.0

// @Router / [get]
func (h *HellowordApi) Add(c *fiber.Ctx) error {
	h.UserService.Add()
	return c.JSON("hello123")
}
