package main

import (
	"fastgo/api/helloword"
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/pkg/app"
	"fastgo/pkg/burst"
	"fastgo/pkg/redoc"
	"fmt"
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func init() {
	logf := "%s[${pid}]%s ${time} ${level} ${path} ${msg}"
	flog.SetFormat(fmt.Sprintf(logf, flog.Cyan, flog.Reset))

	configs.Setup("./configs/helloword.yml")
	conf.Inject()
}

func main() {
	r := app.New(fiber.Config{
		Prefork:     true,
		ReadTimeout: time.Second * 10,
	})
	r.Use(burst.New(burst.Config{
		Limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 20),
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

	r.Install(&helloword.Api{})

	r.Listen(":8080")
}
