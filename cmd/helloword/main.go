package main

import (
	"fastgo2/api/helloword"
	"fastgo2/configs"
	"fastgo2/pkg/grace"
	"fastgo2/pkg/redoc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	configs.Setup("./configs/helloword.yml")
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:     false,
		ReadTimeout: time.Second * 10,
	})
	app.Use(logger.New(logger.Config{
		Format:     "[Fiber] ${time} |${black}${status}|${latency}|${blue}${method} ${url}\n",
		TimeFormat: "2006/01/02 - 15:04:05",
	}))
	app.Get("/doc/*", redoc.New("./api/helloword/swagger.json"))
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
	helloword.Install(app)
	grace.Listen(app, ":8080")
}
