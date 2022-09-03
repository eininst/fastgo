package v1

import (
	"fastgo/internal/service/user"
	"fastgo/pkg/ioc"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func init() {
	ioc.Provide(&HellowordApi{})
}

type HellowordApi struct {
	UserService user.UserService `inject:""`
}

// @Param request body v1.request true "query params"
// @Success 200 {object} v1.HellowordApi.response
// @Router /test [post]
func (h *HellowordApi) Add(c *fiber.Ctx) error {
	fmt.Println(h.UserService)
	h.UserService.Add()
	return c.JSON("hello")
}
