package main

import (
	"sync"
	"time"

	"github.com/domego/gokits/log"
	"github.com/swxctx/bike/bikeWorker/tcpserver"
)

var (
	wg sync.WaitGroup
)

const (
	// listenIp = "192.168.43.50"
	listenIp   = "192.168.0.112"
	listenPort = "8891"
)

func handleApp() {

	// 连接发现服务
	go func() {
		for {
			log.Infof("Listen server start success...")
			tcpserver.DoTcpServer(listenIp, listenPort)
			time.Sleep(time.Second * time.Duration(1))
		}
		wg.Done()
	}()
	wg.Add(1)

	// 数据推送服务
	go func() {
		for {
			log.Infof("Push server start success...")
			tcpserver.WriteConn()
			time.Sleep(time.Second * time.Duration(1))
		}
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
}
