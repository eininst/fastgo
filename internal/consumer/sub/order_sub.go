package sub

import (
	"fastgo/pkg/ioc"
	"github.com/eininst/flog"
	"github.com/eininst/rs"
	"github.com/go-redis/redis/v8"
)

func init() {
	ioc.Provide(&OrderSub{})
}

type OrderSub struct {
	RedisClient *redis.Client `inject:""`
}

func (orderSub *OrderSub) OrderChange(ctx *rs.Context) {
	defer ctx.Ack()
	flog.Info(orderSub.RedisClient)
	flog.Info("abc", ctx.Msg)
}
