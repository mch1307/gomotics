package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/log"
	"github.com/mch1307/gomotics/server"
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
		fmt.Println("conf: ", conf)
		fmt.Println(err)
		panic("Invalid config file/path: file not found ")
	}
	config.Initialize(conf)
	log.Init()
	//log.Info("Starting gomotics")
	s := server.Server{}
	s.Initialize()
	s.Run()
}
