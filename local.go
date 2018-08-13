package lightsocks

import (
	"log"
)

var remoteAddr string

// 新建一个本地端
// 本地端的职责是:
// 1. 监听来自本机浏览器的代理请求
// 2. 转发前加密数据
// 3. 转发socket数据到墙外代理服务端
// 4. 把服务端返回的数据转发给用户的浏览器
// 本地端启动监听，接收来自本机浏览器的连接
func ListenLocal(password string, listenAddr, httpProxyAddr string) error {
	bsPassword, err := parsePassword(password)
	if err != nil {
		return err
	}
	cipher := newCipher(bsPassword)
	remoteAddr = httpProxyAddr
	return ListenSecureTCP(listenAddr, cipher, handleLocalConn)
}

func handleLocalConn(userConn *SecureTCPConn) {
	defer userConn.Close()

	proxyServer, err := DialTCPSecure(remoteAddr, userConn.Cipher)
	if err != nil {
		log.Println(err)
		return
	}
	defer proxyServer.Close()

	// 进行转发
	// 从 proxyServer 读取数据发送到 localUser
	go func() {
		err := proxyServer.DecodeCopy(userConn)
		if err != nil {
			// 在 copy 的过程中可能会存在网络超时等 error 被 return，只要有一个发生了错误就退出本次工作
			userConn.Close()
			proxyServer.Close()
		}
	}()
	// 从 localUser 发送数据发送到 proxyServer，这里因为处在翻墙阶段出现网络错误的概率更大
	userConn.EncodeCopy(proxyServer)
}
