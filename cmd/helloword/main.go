package main

import (
	"fastgo/api/helloword"
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/pkg/burst"
	"fastgo/pkg/di"
	"fastgo/pkg/grace"
	"fastgo/pkg/redoc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang.org/x/time/rate"
	"time"
)

func init() {
	configs.Setup("./configs/helloword.yml")
	conf.Inject()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:     true,
		ReadTimeout: time.Second * 10,
	})
	app.Use(logger.New(logger.Config{
		Format:     "[Fiber] ${time} |${black}${status}|${latency}|${blue}${method} ${url}\n",
		TimeFormat: "2006/01/02 - 15:04:05",
	}))

	app.Use(burst.New(burst.Config{
		Limiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 20),
		Timeout: time.Second * 5,
	}))
	app.Get("/doc/*", redoc.New("./docs/helloword_swagger.json"))

	router := &helloword.Router{Router: app}
	di.Inject(router)
	di.Populate()

	router.Register()
	grace.Listen(app, ":8080")
}
