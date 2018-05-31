package tcpserver

import (
	"net"
	"time"

	"github.com/domego/gokits/log"
)

func DoTcpServer() {
	// 建立socket端口监听
	netListen, err := net.Listen("tcp", "localhost:8891")
	if err != nil {
		log.Errorf("listent port err->%v", err)
	}

	defer netListen.Close()

	log.Infof("Waiting for clients ...")

	//等待客户端访问
	for {
		conn, err := netListen.Accept() //监听接收
		if err != nil {
			continue //如果发生错误，继续下一个循环。
		}

		log.Infof("tcp connect success->IP: %s", conn.RemoteAddr().String()) //tcp连接成功
		go handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {
	// 建立一个slice
	buffer := make([]byte, 2048)
	for {
		// 读取客户端传来的内容
		n, err := conn.Read(buffer)
		if err != nil {
			log.Infof("%s断开连接: %v", conn.RemoteAddr().String())
			// 当远程客户端连接发生错误（断开）后，终止此协程。
			return
		}

		log.Infof("消息来自:%s", conn.RemoteAddr().String())
		log.Infof("消息内容:%s", string(buffer[:n]))

		//返回给客户端的信息
		strTemp := "CofoxServer got msg \"" + string(buffer[:n]) + "\" at " + time.Now().String()
		conn.Write([]byte(strTemp))
	}
}
