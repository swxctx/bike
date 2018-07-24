package tcpserver

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/swxctx/bike/bikeWorker/mail"
	"github.com/swxctx/bike/model/bike"

	"github.com/domego/gokits/log"
)

var (
	connmaps map[string]net.Conn
	Chconn   = make(chan net.Conn, 1024)
	ChMsg    = make(chan string, 1024)
)

func DoTcpServer(ip, port string) {
	connmaps = make(map[string]net.Conn)
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
			continue // 如果发生错误，继续下一个循环。
		}
		thisIp := conn.RemoteAddr().String()
		if _, ok := connmaps[thisIp]; !ok {
			// 不存在该连接
			connmaps[thisIp] = conn
		}
		log.Infof("conns:%v", connmaps)
		log.Infof("tcp connect success->IP: %s", conn.RemoteAddr().String()) //tcp连接成功
		go handleConnection(conn)
	}
}

// 处理连接
func handleConnection(conn net.Conn) {
	// 建立一个slice
	buffer := make([]byte, 2048)
	for {
		// 读取客户端传来的内容
		n, err := conn.Read(buffer)
		if err != nil {
			log.Infof("%s断开连接", conn.RemoteAddr().String())
			// 当远程客户端连接发生错误（断开）后，终止此协程。
			return
		}

		log.Infof("消息来自:%s", conn.RemoteAddr().String())
		log.Infof("消息内容:%s", string(buffer[:n]))

		ip := conn.RemoteAddr().String()
		message := string(buffer[:n])
		if !strings.Contains(message, "AT") && !strings.Contains(message, "msg") {
			go DoDealData(ip, message)
		}
		// 处理推送服务(单车开锁服务)
		if strings.Contains(message, "msg") {
			msg := unmarShalPushInfo(message)
			pushMap := connmaps[msg.Ip]

			Chconn <- pushMap
			ChMsg <- msg.Msg
			log.Infof("push bike addr:%s", pushMap.RemoteAddr().String())
		}
		//返回给客户端的信息(延迟3s进行，心跳发送)
		time.Sleep(time.Second * time.Duration(3))

		if len(message) > 0 {
			ChMsg <- "11"
			Chconn <- conn
		}
	}
}

// WriteConn 数据响应服务
func WriteConn() {
	for {
		msg := <-ChMsg
		conn := <-Chconn
		log.Infof("push msg :%s", msg)
		log.Infof("push Ip:%s", conn.RemoteAddr().String())
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Errorf("write err->%v", err)
		}
		log.Infof("recive send...")
		//返回给客户端的信息(延迟3s进行，心跳发送)
		time.Sleep(time.Second * time.Duration(2))
	}
}

// DoDealData 处理接收到的数据
func DoDealData(ip, message string) {
	var (
		port string
	)
	ipinfo := strings.Split(ip, ":")
	if len(ipinfo) > 1 {
		ip = ipinfo[0]
		port = ipinfo[1]
	}

	// 查询单车编号
	bikeInfo := bike.GetBikeFirst("`deleted`=0 AND `ip`=?", ip)
	if bikeInfo == nil {
		log.Errorf("This addr bike is nil->addr:%s", ip)
		return
	}
	// 更新TCP端口
	bikeInfo.Port = port
	bikeInfo.Update()
	// 判断数据是否异常
	if len(message) == 0 {
		log.Errorf("message is nil->addr:%s", ip)
		return
	}
	msg := strings.Split(message, "|")
	log.Infof("msg is: %v", msg)
	if msg[0] == "0" && msg[1] == "0" && msg[2] == "1" {
		// 数据异常，存储并且进行报警
		mail.SendMail(bikeInfo.ById)
		bikeAlarm := bike.NewBikeAlarm()
		bikeAlarm.AlarmTs = time.Now().Unix()
		bikeAlarm.ById = bikeInfo.ById
		mh1, _ := strconv.ParseInt(msg[0], 10, 64)
		bikeAlarm.MhCountOne = int32(mh1)
		mh2, _ := strconv.ParseInt(msg[1], 10, 64)
		bikeAlarm.MhCountTwo = int32(mh2)
		sw, _ := strconv.ParseInt(msg[2], 10, 64)
		bikeAlarm.SwCountOne = int32(sw)
		bikeAlarm.Insert()
	}
}

func unmarShalPushInfo(data string) *PushInfo {
	r := &PushInfo{}
	if err := json.Unmarshal([]byte(data), r); err != nil {
		log.Errorf("resource json unmarshal error: %v", err)
	}
	return r
}

type PushInfo struct {
	Ip  string `json:"ip"`
	Msg string `json:"msg"`
}
