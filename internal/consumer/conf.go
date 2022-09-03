package consumer

import (
	"fastgo/internal/consumer/sub"
	"github.com/eininst/rs"
)

type Conf struct {
	rs.Client
	OrderSub *sub.OrderSub `inject:""`
}

func (c *Conf) Install() {
	c.Receive(rs.Rctx{Stream: "test", Handler: c.OrderSub.OrderChange})
}
