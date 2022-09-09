package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	App       *fiber.App       `inject:""`
	Helloword *v1.HellowordApi `inject:""`
}

func (api *Api) Init() {
	api.App.Get("/add", api.Helloword.Add)
	api.App.Post("/user", v1.AddUser)
}
