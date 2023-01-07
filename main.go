package main

import (
	"flag"

	"github.com/YogeLiu/CloudDisk/bootstrap"
	"github.com/YogeLiu/CloudDisk/pkg/conf"
	"github.com/YogeLiu/CloudDisk/routers"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", "", "配置文件")
	flag.Parse()
	bootstrap.Init(confPath)
}

func main() {
	engine := routers.InitRouter()
	engine.Run(conf.SystemConfig.Listen)
}
