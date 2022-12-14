package consumer

import (
	"fastgo/consumer/sub"
	"github.com/eininst/rs"
)

type Conf struct {
	Cli      rs.Client     `inject:""`
	OrderSub *sub.OrderSub `inject:""`
}

func (f *Conf) Init() {
	f.Cli.Receive(rs.Rctx{
		Stream:  "test",
		Handler: f.OrderSub.OrderChange,
	})
}
