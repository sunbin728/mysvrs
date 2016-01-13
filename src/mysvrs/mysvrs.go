package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"mysvrs/config"
	"mysvrs/myweb"
	"mysvrs/util"

	"github.com/nybuxtsui/log"
)

var (
	servconf = flag.String("servconf", "mysvrs.conf", "服务配置文件")
	logconf  = flag.String("logconf", "logger.conf", "日志配置文件")
)

func main() {
	flag.Parse()
	gomax := int(float64(runtime.NumCPU()) * 1.5)
	runtime.GOMAXPROCS(gomax)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	log.InitFromFile(*logconf)
	log.Info("version:%v", util.GetVersion())
	log.Info("start...")

	Init()

	go myweb.Start(config.SConfig.Http, config.SConfig.Https)
	log.Info("Welcome to mysvrs")

	<-signalChan
	log.Info("bye")
	os.Exit(0)

}

func Init() {
	config.Init(*servconf)
	myweb.Init()
}
