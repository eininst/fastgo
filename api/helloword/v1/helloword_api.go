package v1

import (
	"fastgo/internal/common/inject"
	"fastgo/internal/service/user"
	"github.com/gofiber/fiber/v2"
)

func init() {
	inject.Provide(new(HellowordApi))
}

type HellowordApi struct {
	UserService user.UserService `inject:""`
}

// @Summary 测试swagger
// @Tags test
// @version 1.0

// @Router / [get]
func (h *HellowordApi) Add(c *fiber.Ctx) error {
	er := h.UserService.Add()
	if er != nil {
		return er
	}
	return c.JSON("hello123")
}
