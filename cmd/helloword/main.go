package main

import (
	"fastgo/api/helloword"
	"fastgo/configs"
	"fastgo/internal/common/inject"
	"fastgo/internal/common/middleware/redoc"
	"fastgo/internal/conf"
	"fmt"
	burst "github.com/eininst/fiber-middleware-burst"
	grace "github.com/eininst/fiber-prefork-grace"
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func init() {
	logf := "%s[${pid}]%s ${time} ${level} ${path} ${msg}"
	flog.SetFormat(fmt.Sprintf(logf, flog.Cyan, flog.Reset))

	configs.SetConfig("./configs/helloword.yml")
	conf.Provide()
}

func main() {
	r := fiber.New(fiber.Config{
		Prefork:     true,
		ReadTimeout: time.Second * 10,
	})
	r.Use(burst.New(burst.Config{
		Limiter: rate.NewLimiter(200, 500),
		Timeout: time.Second * 5,
	}))

	r.Use(logger.New(logger.Config{
		Format:     "[Fiber] [${pid}] ${time} |${black}${status}|${latency}|${blue}${method} ${url}\n",
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	r.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	r.Get("/doc/*", redoc.New("api/helloword/swagger.json"))

	r.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	r.Use(recover.New())

	var api helloword.Api
	inject.Provide(&api)
	inject.Populate()
	api.Router(r)

	grace.Listen(r, "8080")
}
