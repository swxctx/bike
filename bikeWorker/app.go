package main

import (
	"sync"

	"github.com/swxctx/bike/bikeWorker/tcpserver"
)

var (
	wg sync.WaitGroup
)

func handleApp() {

	go func() {
		for {
			tcpserver.DoTcpServer()
		}
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
}
