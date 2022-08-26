package burst

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
	"time"
)

type Config struct {
	Limit          rate.Limit
	Burst          int
	Timeout        time.Duration
	TimeoutHandler fiber.Handler
}

func New(cfg Config) fiber.Handler {
	if cfg.Limit == 0 {
		panic("Limit cannot be empty")
	}
	if cfg.Burst == 0 {
		panic("Burst cannot be empty")
	}
	if cfg.Timeout == 0 {
		panic("Timeout cannot be empty")
	}

	burstLimit := rate.NewLimiter(cfg.Limit, cfg.Burst)

	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), cfg.Timeout)
		defer cancel()
		err := burstLimit.Wait(ctx)
		if err != nil {
			if cfg.TimeoutHandler == nil {
				return c.SendStatus(fiber.StatusTooManyRequests)
			}
			return cfg.TimeoutHandler(c)
		}
		return c.Next()
	}
}
