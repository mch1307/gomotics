package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

// ServerConf holds the server config
type ServerConf struct {
	ListenPort int    `toml:"listenPort"`
	LogLevel   string `toml:"logLevel"`
	LogPath    string `toml:"logPath"`
	GMHostPort string `toml:"gmHostPort"`
}

// JeedomConf holds the server config
type JeedomConf struct {
	URL               string `toml:"url"`
	APIKey            string `toml:"apikey"`
	Enabled           bool
	AutoCreateObjects bool `toml:"autoCreateObjects"`
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

// coalesce returns the first non nil/zero passed value as string
// numeric 0 value considered as empty
func coalesce(data ...interface{}) string {
	for _, v := range data {
		switch v := v.(type) {
		case string:
			if len(v) > 0 {
				return v
			}
		case int:
			if v > 0 {
				strconv.Itoa(v)
				return strconv.Itoa(v)
			}
		}
	}
	return ""
}

// Initialize populates the Conf variable
func Initialize(cfg string) {
	Conf.JeedomConfig.Enabled = false
	// load config file if any
	if cfg != "" {
		if _, err := os.Stat(cfg); err != nil {
			wrkDir, _ := os.Getwd()
			Conf.ServerConfig.LogPath = wrkDir
		} else {
			if _, err := toml.DecodeFile(cfg, &Conf); err != nil {
				fmt.Println("Error parsing config file: ", err)
			}
		}
	}
	wrkDir, _ := os.Getwd()
	listenPort, _ := strconv.Atoi(coalesce(os.Getenv("LISTEN_PORT"), Conf.ServerConfig.ListenPort, "8081"))
	Conf.ServerConfig.ListenPort = listenPort
	Conf.ServerConfig.LogLevel = coalesce(os.Getenv("LOG_LEVEL"), Conf.ServerConfig.LogLevel, "INFO")
	Conf.ServerConfig.LogPath = coalesce(os.Getenv("LOG_PATH"), Conf.ServerConfig.LogPath, wrkDir)
	Conf.JeedomConfig.URL = coalesce(os.Getenv("JEE_URL"), Conf.JeedomConfig.URL)
	Conf.JeedomConfig.APIKey = coalesce(os.Getenv("JEE_APIKEY"), Conf.JeedomConfig.APIKey)
	Conf.NhcConfig.Host = coalesce(os.Getenv("NHC_HOST"), Conf.NhcConfig.Host)
	Conf.ServerConfig.GMHostPort = coalesce(os.Getenv("GM_HOSTPORT"), Conf.ServerConfig.GMHostPort, "localhost:"+strconv.Itoa(Conf.ServerConfig.ListenPort))
	var co = "N"
	if Conf.JeedomConfig.AutoCreateObjects {
		co = "Y"
	}
	co = coalesce(os.Getenv("AUTO_CREATE_OBJECTS"), co, "N")
	if strings.ToUpper(co) == "Y" {
		Conf.JeedomConfig.AutoCreateObjects = true
	}

	nhcPort, _ := strconv.Atoi(coalesce(os.Getenv("NHC_PORT"), Conf.NhcConfig.Port, "8000"))
	Conf.NhcConfig.Port = nhcPort

	if len(Conf.JeedomConfig.APIKey) > 0 {
		Conf.JeedomConfig.Enabled = true
	}
	fmt.Printf("Starting with config: %+v", Conf)
	fmt.Println(" ")
}
