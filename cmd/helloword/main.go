package main

import (
	"fastgo/api/helloword"
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/pkg/app"
	"fmt"
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2"
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
	r.Install(&helloword.Api{})

	r.Listen(":8080")
}
