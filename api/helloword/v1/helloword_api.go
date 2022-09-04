package v1

import (
	"fastgo/internal/service/user"
	"fastgo/pkg/di"
	"github.com/gofiber/fiber/v2"
)

func init() {
	di.Inject(&HellowordApi{})
}

type HellowordApi struct {
	UserService user.UserService `inject:""`
}

// @Router /test [post]
func (h *HellowordApi) Add(c *fiber.Ctx) error {
	h.UserService.Add()
	return c.JSON("hello123")
}
