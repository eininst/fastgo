package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"fastgo/pkg/redoc"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Api struct {
	Helloword *v1.HellowordApi `inject:""`
}

func (api *Api) Router(router fiber.Router) {
	router.Get("/doc/*", redoc.New("docs/helloword_swagger.json"))

	router.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	router.Get("/accounts/:id", api.Helloword.Add)
}
