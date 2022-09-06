package helloword

import (
	v1 "fastgo/api/helloword/v1"
	"fastgo/pkg/burst"
	"fastgo/pkg/redoc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type Api struct {
	Helloword *v1.HellowordApi `inject:""`
}

func (api *Api) Router(r fiber.Router) {
	r.Use(burst.New(burst.Config{
		Limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 20),
		Timeout: time.Second * 5,
	}))

	r.Use(logger.New(logger.Config{
		Format:     "[Fiber] [${pid}] ${time} |${black}${status}|${latency}|${blue}${method} ${url}\n",
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	r.Get("/doc/*", redoc.New("api/helloword/swagger.json"))

	r.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	r.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	r.Get("/accounts/:id", api.Helloword.Add)
}
