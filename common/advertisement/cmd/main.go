package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"easymarket-go-server/common/advertisement/internal/di"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

func initLog() {
	var (
		dc struct {
			Log *log.Config
		}
	)
	if err := paladin.Get("log.toml").UnmarshalTOML(&dc); err != nil {
		panic(err)
	}
	log.Init(dc.Log)
}
func main() {
	flag.Parse()
	paladin.Init()
	initLog()
	defer log.Close()
	log.Info("goods start")
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("goods exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
