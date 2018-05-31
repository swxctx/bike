package tcpserver

import (
	"net"

	"github.com/domego/gokits/log"
)

func DoTcpClient(server string, message string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		log.Errorf("ResolveTCPAddr err-> %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Errorf("DialTCP err-> %v", err)
	}

	log.Infof("connection success")
	sender(conn, message)
}

//发送信息
func sender(conn net.Conn, message string) {
	conn.Write([]byte(message))
	log.Infof("send over")

	//接收服务端反馈
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Errorf("waiting server back msg err->%v,Ip->%s", err, conn.RemoteAddr().String())
		return
	}
	log.Infof("%s->receive server back msg: %s", conn.RemoteAddr().String(), string(buffer[:n]))
}
