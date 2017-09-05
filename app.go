package main

import (
	"flag"

	"github.com/mch1307/gomotics/log"

	"github.com/mch1307/gomotics/config"
	"github.com/mch1307/gomotics/nhc"
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
	Sub(conf)
}

// Sub actually starts the servers
func Sub(conf string) {
	config.Initialize(conf)
	log.Init()
	s := server.Server{}
	s.Initialize()
	s.Run()
	nhc.Init(&config.Conf.NhcConfig)
}
