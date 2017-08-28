package main

import (
	"flag"
	"fmt"
	"os"

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
		fmt.Println("Invalid config file/path, file not found: ", err)
		panic(err)
	}
	s := server.Server{}
	s.Initialize(conf)
	s.Run()
}
