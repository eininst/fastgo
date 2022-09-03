package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"fastgo/pkg/di"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiber.Router
	HellowordApi *v1.HellowordApi `inject:""`
}

func NewRouter(r fiber.Router) *Router {
	router := &Router{Router: r}
	di.Inject(router)
	di.Populate()
	return router
}

func (r *Router) Register() {
	r.Get("/accounts/:id", r.HellowordApi.Add)
}
