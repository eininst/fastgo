package main

import (
	"fastgo/configs"
	"fastgo/consumer"
	"fastgo/internal/conf"
	"github.com/eininst/flog"
	"github.com/eininst/ninja"
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
	ninja.Install(&c)

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		c.Cli.Shutdown()
		flog.Info("Graceful Shutdown")
	}()

	c.Cli.Listen()
}
