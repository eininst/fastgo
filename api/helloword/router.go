// @title          Fiber Example API
// @version        1.0
// @description    This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  fiber@swagger.io
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           localhost:8080
// @x-example-key {"key": "value"}
// @BasePath       /
package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
	HellowordApi *v1.HellowordApi `inject:""`
}

func (r *Router) Register() {
	r.Get("/hello", r.HellowordApi.Add)
}
