package api

import (
	"fastgo/configs"
	burst "github.com/eininst/fiber-middleware-burst"
	recovers "github.com/eininst/fiber-middleware-recover"
	redoc "github.com/eininst/fiber-middleware-redoc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/tidwall/gjson"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func init() {
	registerDefault()
}

var (
	defaultName = "middleware"
	mux         = &sync.Mutex{}
)

type MiddlewareHandler func(app *fiber.App, value gjson.Result)

var handlerx = map[string]MiddlewareHandler{}

func Register(name string, handler MiddlewareHandler) {
	mux.Lock()
	defer mux.Unlock()
	handlerx[name] = handler
}

func CommonMiddleware(app *fiber.App, cfg ...gjson.Result) {
	var res gjson.Result
	if len(cfg) > 0 {
		res = cfg[0]
	} else {
		res = configs.Get(defaultName)
	}

	for _, r := range res.Array() {
		r.ForEach(func(key, value gjson.Result) bool {
			if handler, ok := handlerx[key.String()]; ok {
				if value.Exists() {
					handler(app, value)
				}
			}
			return true
		})
	}
}

func registerDefault() {
	Register("recover", func(app *fiber.App, value gjson.Result) {
		rdefaltCfg := recovers.Config{}
		stackBuflenRes := value.Get("stackBuflen")
		if stackBuflenRes.Exists() {
			rdefaltCfg.StackTraceBufLen = int(stackBuflenRes.Int())
		}
		app.Use(recovers.New(rdefaltCfg))
	})

	Register("limiter", func(app *fiber.App, value gjson.Result) {
		rt := value.Get("rate").Int()
		bst := value.Get("burst").Int()
		timeout := value.Get("timeout").Int()
		if rt != 0 && bst != 0 && timeout != 0 {
			app.Use(burst.New(burst.Config{
				Limiter: rate.NewLimiter(rate.Limit(rt), int(bst)),
				Timeout: time.Second * time.Duration(timeout),
			}))
		}
	})

	Register("ready", func(app *fiber.App, value gjson.Result) {
		path := "/ready"
		pathRes := value.Get("path")
		if pathRes.String() != "" {
			path = pathRes.String()
		}
		app.Get(path, func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(http.StatusOK)
		})
	})

	Register("logger", func(app *fiber.App, value gjson.Result) {
		f := "[Fiber] [${pid}] ${time} |${black}${status}|${latency}|${blue}${method} ${url}\n"
		tf := "2006/01/02 - 15:04:05"
		if value.Get("format").String() != "" {
			f = value.Get("format").String()
		}
		if value.Get("timeFormat").String() != "" {
			f = value.Get("timeFormat").String()
		}

		app.Use(logger.New(logger.Config{
			Format:     f,
			TimeFormat: tf,
		}))
	})

	Register("monitor", func(app *fiber.App, value gjson.Result) {
		app.Get("/metrics", monitor.New())
	})

	Register("swagger", func(app *fiber.App, value gjson.Result) {
		path := value.Get("path").String()
		json := value.Get("json").String()
		if path != "" && json != "" {
			app.Get(path, redoc.New(json))
		}
	})
}
