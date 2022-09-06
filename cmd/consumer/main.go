package main

import (
	"fastgo/configs"
	"fastgo/internal/conf"
	"fastgo/internal/consumer"
	"fastgo/pkg/di"
	"github.com/eininst/flog"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configs.Setup("./configs/consumer.yml")
	conf.Inject()
}

func main() {
	var c consumer.Conf
	di.Inject(&c)
	di.Populate()

	c.Subscribe()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		c.Cli.Shutdown()
		flog.Info("Graceful Shutdown")
	}()

	c.Cli.Listen()
}
