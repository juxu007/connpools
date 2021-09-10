package connpools

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego/logs"
)

func turnOnGracefulExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for s := range ch {
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT:
				logs.Info("exit on signal: ", s)
				gracefulExit()
			default:
				logs.Info("get signal: ", s)
			}
		}
	}()
}
