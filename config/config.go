package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// ServerConf holds the server config
type ServerConf struct {
	ListenPort int    `toml:"ListenPort"`
	LogLevel   string `toml:"LogLevel"`
	LogPath    string `toml:"LogPath"`
}

// JeedomConf holds the server config
type JeedomConf struct {
	Service string
	Host    string
	Port    int
	APIPath string
	APIKey  string
}

// NhcConf holds the server config
type NhcConf struct {
	Host string
	Port int
}

// GlobalConfig holds the structure of the yml configuration file
// It has 3 sections: server, jeedom and nhc
type GlobalConfig struct {
	ServerConfig ServerConf `toml:"server"`
	JeedomConfig JeedomConf `toml:"jeedom"`
	NhcConfig    NhcConf    `toml:"nhc"`
}

// Conf holds the global configuration
var Conf GlobalConfig

// Initialize populates the Conf variable
func Initialize(cfg string) error {
	var err error
	if _, err := os.Stat(cfg); err != nil {
		fmt.Println("Invalid config file/path: ", err)
		wrkDir, _ := os.Getwd()
		Conf.ServerConfig.ListenPort = 8081
		Conf.ServerConfig.LogLevel = "INFO"
		Conf.ServerConfig.LogPath = wrkDir
		fmt.Printf("Starting with default config: %+v", Conf.ServerConfig)
		fmt.Println(" ")
	} else {
		if _, err := toml.DecodeFile(cfg, &Conf); err != nil {
			return err
		}
	}
	return err
}
