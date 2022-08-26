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
	v1 "fastgo2/api/helloword/v1"
	"github.com/gofiber/fiber/v2"
)

func Install(app *fiber.App) {
	app.Get("/hello", v1.HelloWorld)
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("2")
	})
}
