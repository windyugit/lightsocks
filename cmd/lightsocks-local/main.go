package main

import (
	"log"

	"github.com/gwuhaolin/lightsocks"
	"github.com/gwuhaolin/lightsocks/cmd"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 默认配置
	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 local 端并监听
	log.Fatalln(lightsocks.ListenLocal(config.Password, config.ListenAddr, config.RemoteAddr))
}
