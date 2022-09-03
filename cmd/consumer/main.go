package main

import (
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/internal/consumer"
	"fastgo/internal/data/rdb"
	"fastgo/pkg/ioc"
	"github.com/eininst/flog"
	"github.com/eininst/rs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	configs.Setup("./configs/consumer.yml")
	conf.Inject()
}

func main() {
	cli := rs.New(rdb.Get(), rs.
		Config{
		Receive: rs.ReceiveConfig{
			Work:       rs.Int(10),
			ReadCount:  rs.Int64(1),
			BlockTime:  time.Second * 20,
			MaxRetries: rs.Int64(3),
			Timeout:    time.Second * 20,
		},
	})
	c := &consumer.Conf{Client: cli}
	ioc.Provide(c)
	ioc.Populate()

	c.Install()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		cli.Shutdown()
		flog.Info("Graceful Shutdown")
	}()

	cli.Listen()
}
