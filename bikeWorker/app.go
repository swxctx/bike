package main

import (
	"sync"

	"github.com/swxctx/bike/bikeWorker/tcpserver"
)

var (
	wg sync.WaitGroup
)

const (
	listenIp   = "127.0.0.1"
	listenPort = "8891"
)

func handleApp() {

	go func() {
		for {
			tcpserver.DoTcpServer(listenIp, listenPort)
		}
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
}
