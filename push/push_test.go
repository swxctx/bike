package push

import (
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	for i := 0; i < 10; i++ {
		DoTcpClient("192.168.0.108:8891", "1")
		time.Sleep(time.Second * time.Duration(2))
	}
}
