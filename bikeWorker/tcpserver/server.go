package tcpserver

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/swxctx/bike/bikeWorker/mail"
	"github.com/swxctx/bike/model/bike"

	"github.com/domego/gokits/log"
)

func DoTcpServer(ip, port string) {
	addr := fmt.Sprintf("%s:%s", ip, port)
	// 建立socket端口监听
	netListen, err := net.Listen("tcp", addr)
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

		ip := conn.RemoteAddr().String()
		meaasge := string(buffer[:n])
		DoDealData(ip, meaasge)
		//返回给客户端的信息
		strTemp := "CofoxServer got msg \"" + string(buffer[:n]) + "\" at " + time.Now().String()
		conn.Write([]byte(strTemp))
	}
}

// DoDealData 处理接收到的数据
func DoDealData(ip, message string) {
	// 查询单车编号
	bikeInfo := bike.GetBikeFirst("`deleted`=0 AND `addr`=?", ip)
	if bikeInfo == nil {
		log.Errorf("This addr bike is nil->addr:%s", ip)
		return
	}
	// 判断数据是否异常
	if len(message) == 0 {
		log.Errorf("message is nil->addr:%s", ip)
		return
	}
	msg := unmarShalMsg(message)
	if msg == nil {
		log.Errorf("unmarshal mesg is nil:%s", ip)
		return
	}
	if msg.MhCountOne == 0 && msg.SwCountOne == 1 && msg.SwCountTwo == 1 {
		// 数据异常，存储并且进行报警
		mail.SendMail(bikeInfo.ById)
		bikeAlarm := bike.NewBikeAlarm()
		bikeAlarm.AlarmTs = time.Now().Unix()
		bikeAlarm.ById = bikeInfo.ById
		bikeAlarm.MhCountOne = msg.MhCountOne
		bikeAlarm.MhCountTwo = msg.MhCountTwo
		bikeAlarm.SwCountOne = msg.SwCountOne
		bikeAlarm.SwCountTwo = msg.SwCountTwo
		bikeAlarm.Insert()
	}
}

func unmarShalMsg(data string) *Message {
	r := &Message{}
	if err := json.Unmarshal([]byte(data), r); err != nil {
		log.Errorf("resource json unmarshal error: %v", err)
	}
	return r
}

type Message struct {
	MhCountOne int32 `json:"mh_count_one"`
	MhCountTwo int32 `json:"mh_count_two"`
	SwCountOne int32 `json:"sw_count_one"`
	SwCountTwo int32 `json:"sw_count_two"`
}
