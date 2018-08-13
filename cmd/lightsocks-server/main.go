package main

import (
	"fmt"
	"log"

	"github.com/gwuhaolin/lightsocks"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/phayes/freeport"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 服务端监听端口随机生成
	port, err := freeport.GetFreePort()
	if err != nil {
		// 随机端口失败就采用 7448
		port = 7448
	}
	// 默认配置
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		// 密码随机生成
		Password: lightsocks.RandPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// 启动 server 端并监听
	log.Fatalln(lightsocks.ListenServer(config.Password, config.ListenAddr))
}
