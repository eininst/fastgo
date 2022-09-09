package main

import (
	"fastgo/api"
	"fastgo/api/helloword"
	"fastgo/configs"
	"fastgo/internal/common/serr"
	"fastgo/internal/conf"
	"fmt"
	grace "github.com/eininst/fiber-prefork-grace"
	"github.com/eininst/flog"
	"github.com/eininst/ninja"
	"github.com/gofiber/fiber/v2"
	"time"
)

func init() {
	logf := "%s[${pid}]%s ${time} ${level} ${path} ${msg}"
	flog.SetFormat(fmt.Sprintf(logf, flog.Cyan, flog.Reset))

	configs.SetConfig("./configs/helloword.yml")
	conf.Provide()
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		ErrorHandler: serr.ErrorHandler,
	})

	api.CommonMiddleware(app)

	var hapi helloword.Api
	ninja.Install(&hapi, app)

	grace.Listen(app, ":8080")
}
