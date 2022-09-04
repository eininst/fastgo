package app

import (
	"fastgo/pkg/di"
	"fastgo/pkg/grace"
	"github.com/gofiber/fiber/v2"
	"time"
)

type App struct {
	*fiber.App
}

type Api interface {
	Router(r fiber.Router)
}

func New(config ...fiber.Config) *App {
	return &App{
		App: fiber.New(config...),
	}
}

func (a *App) Install(api Api) {
	di.Inject(api)
	di.Populate()
	api.Router(a.App)
}

func (a *App) Listen(addr string) {
	grace.Listen(a.App, addr)
}

func (a *App) ListenTLS(addr, certFile, keyFile string, timeout ...time.Duration) {
	grace.ListenTLS(a.App, addr, certFile, keyFile, timeout...)
}
