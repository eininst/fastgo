package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiber.Router
	HellowordApi *v1.HellowordApi `inject:""`
}

func (r *Router) Register() {

	r.Get("/accounts/:id", r.HellowordApi.Add)
}
