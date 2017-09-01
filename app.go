package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/server"
	"github.com/mch1307/gomotics/ws"
)

var (
	conf string
)

func init() {
	flag.StringVar(&conf, "conf", "", "Config file name (full path to TOML file)")
}

func main() {
	flag.Parse()
	if _, err := os.Stat(conf); err != nil {
		fmt.Println("Invalid config file/path, file not found: ", err)
		panic(err)
	}
	Sub(conf)
}

// Sub actually starts the servers
func Sub(conf string) {
	config.Initialize(conf)
	s := server.Server{}
	s.Initialize()
	s.Run()
	ws.Initialize()
}
