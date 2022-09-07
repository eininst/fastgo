package main

import (
	"fastgo/common/inject"
	"fastgo/configs"
	"fastgo/consumer"
	"fastgo/internal/conf"
	"github.com/eininst/flog"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configs.SetConfig("./configs/consumer.yml")
	conf.Provide()
}

func main() {
	var c consumer.Conf
	inject.Provide(&c)
	inject.Populate()

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
