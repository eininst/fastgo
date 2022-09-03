package main

import (
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/internal/consumer"
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
	c := consumer.New(rs.Config{
		Receive: rs.ReceiveConfig{
			Work:       rs.Int(10),
			ReadCount:  rs.Int64(1),
			BlockTime:  time.Second * 20,
			MaxRetries: rs.Int64(3),
			Timeout:    time.Second * 20,
		},
	})
	c.Subscribe()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		c.Client.Shutdown()
		flog.Info("Graceful Shutdown")
	}()

	c.Client.Listen()
}
