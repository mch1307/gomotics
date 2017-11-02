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
	URL     string `toml:"url"`
	APIKey  string `toml:"apikey"`
	Enabled bool
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
func Initialize(cfg string) {

	if _, err := os.Stat(cfg); err != nil {
		//fmt.Println("Invalid config file/path: ", err)
		wrkDir, _ := os.Getwd()
		Conf.ServerConfig.LogPath = wrkDir
	} else {
		if _, err := toml.DecodeFile(cfg, &Conf); err != nil {
			fmt.Println("Error parsing config file: ", err)
		}
	}
	if Conf.ServerConfig.ListenPort == 0 {
		Conf.ServerConfig.ListenPort = 8081
	}
	if len(Conf.ServerConfig.LogLevel) == 0 {
		Conf.ServerConfig.LogLevel = "INFO"
	}
	if len(Conf.JeedomConfig.APIKey) > 0 {
		Conf.JeedomConfig.Enabled = true
	}
	fmt.Printf("Starting with config: %+v", Conf.ServerConfig)
	fmt.Println(" ")
}
