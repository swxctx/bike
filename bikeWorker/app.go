package main

import (
	"sync"
	"time"

	"github.com/swxctx/bike/bikeWorker/mail"
	"github.com/swxctx/bike/bikeWorker/tcpserver"
)

var (
	wg sync.WaitGroup
)

func handleApp() {

	go func() {
		for {
			mail.SendMail()
			time.Sleep(time.Second * time.Duration(100))
			tcpserver.DoTcpServer()
		}
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
}
