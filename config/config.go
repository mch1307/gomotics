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
	Host         string
	Port         int
	RegisterCmd  string
	GetEquipCmd  string
	GetLocCmd    string
	GetEnergyCmd string
	GetThermoCmd string
}

// GlobalConfig holds the structure of the yml configuration file
// It has 3 sections: server, jeedom and nhc
type GlobalConfig struct {
	ServerConfig ServerConf `toml:"server"`
	JeedomConfig JeedomConf `toml:"jeedom"`
	NhcConfig    NhcConf    `toml:"nhc"`
}

// GetConf return the global app configuration from toml file
func GetConf() (config GlobalConfig, err error) {
	var cfg GlobalConfig
	if _, err := toml.DecodeFile("./config/config.toml", &cfg); err != nil {
		fmt.Println(err)
	}
	return cfg, err
}
