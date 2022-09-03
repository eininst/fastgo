package consumer

import (
	"fastgo/internal/consumer/sub"
	"fastgo/pkg/di"
	"github.com/eininst/rs"
	"github.com/go-redis/redis/v8"
)

type Conf struct {
	rs.Client
	RedisClient *redis.Client `inject:""`
	OrderSub    *sub.OrderSub `inject:""`
}

func New(rsConf rs.Config) *Conf {
	cf := &Conf{}
	di.Inject(cf)
	di.Populate()

	cli := rs.New(cf.RedisClient, rsConf)
	cf.Client = cli
	return cf
}

func (c *Conf) Subscribe() {
	c.Receive(rs.Rctx{Stream: "test", Handler: c.OrderSub.OrderChange})
}
