package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// ServerConf holds the server config
type ServerConf struct {
	ListenPort int
	LogLevel   string
	LogPath    string
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
	if _, err := toml.DecodeFile(cfg, &Conf); err != nil {
		fmt.Println(err)
	}
	return err
}
