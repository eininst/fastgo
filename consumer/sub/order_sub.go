package sub

import (
	"fastgo/internal/common/inject"
	"github.com/eininst/flog"
	"github.com/eininst/rs"
	"github.com/go-redis/redis/v8"
)

func init() {
	inject.Provide(&OrderSub{})
}

type OrderSub struct {
	RedisClient *redis.Client `inject:""`
}

func (orderSub *OrderSub) OrderChange(ctx *rs.Context) {
	defer ctx.Ack()

	flog.Info(orderSub.RedisClient)
	flog.Info("abc", ctx.Msg)
}
