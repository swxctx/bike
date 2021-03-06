package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"

  "github.com/domego/gokits/log"
  "github.com/swxctx/bike/bikeApi/cfg"
)

var (
	conf     *cfg.Config
	appName string
	pid     int
)

var (
	pidFile     = "log/app.pid"
	appNameFile = "NAME"
)

func init() {
	var _ = os.Mkdir("log", os.ModePerm)
	bs, err := ioutil.ReadFile(appNameFile)
	if err != nil {
		log.Errorf("read appName error, the path is %s, error: %s", appNameFile, err)
	}
	appName = strings.Trim(string(bs), "\n")
}

func writePid() {
	pid = syscall.Getpid()
	ioutil.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

func handleEnd() {
  // os.Remove(pidFile)
}

func handleReloadSignal() {
  cfg.Reload()
}

func main() {
	flag.Parse()
	writePid()

  conf = cfg.Init()
	defer log.Infof("end")
	handleApp()
}
