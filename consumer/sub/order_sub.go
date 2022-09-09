package sub

import (
	"github.com/eininst/flog"
	"github.com/eininst/ninja"
	"github.com/eininst/rs"
	"github.com/go-redis/redis/v8"
)

func init() {
	ninja.Provide(&OrderSub{})
}

type OrderSub struct {
	RedisClient *redis.Client `inject:""`
}

func (orderSub *OrderSub) OrderChange(ctx *rs.Context) {
	defer ctx.Ack()

	flog.Info(orderSub.RedisClient)
	flog.Info("abc", ctx.Msg)
}
