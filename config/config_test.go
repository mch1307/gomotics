package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	//. "github.com/mch1307/gomotics/config"
)

func init() {
	fmt.Println("starting config test")
}
func TestInitialize(t *testing.T) {
	var cfg GlobalConfig
	var cfgDefault GlobalConfig
	cfg.NhcConfig.Host = "localhost"
	cfg.NhcConfig.Port = 8000
	cfg.ServerConfig.ListenPort = 8081
	cfg.ServerConfig.LogLevel = "DEBUG"
	cfg.ServerConfig.LogPath = "."
	cfg.JeedomConfig.URL = "http://jeedom/core/api/jeeApi.php"
	cfg.JeedomConfig.APIKey = ""
	cfg.JeedomConfig.Enabled = false
	cfg.JeedomConfig.CreateNHCObjects = false
	cfg.ServerConfig.EndpointURL = "localhost"
	cfgDefault.NhcConfig.Host = ""
	cfgDefault.NhcConfig.Port = 8000
	cfgDefault.ServerConfig.ListenPort = 8081
	cfgDefault.ServerConfig.LogLevel = "INFO"
	cfgDefault.ServerConfig.LogPath, _ = os.Getwd()
	cfgDefault.JeedomConfig.URL = ""
	cfgDefault.JeedomConfig.APIKey = ""
	cfgDefault.JeedomConfig.Enabled = false
	cfgDefault.JeedomConfig.CreateNHCObjects = false
	cfgDefault.ServerConfig.EndpointURL = "localhost"
	tests := []struct {
		name     string
		confFile string
		wantErr  bool
		exCcfg   GlobalConfig
	}{
		{"read conf - listen port", "test.toml", false, cfg},
		{"read conf - nhc host", "test.toml", false, cfg},
		{"read conf - no file", "NoSuchFile", false, cfgDefault},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize(tt.confFile)
			if (!tt.wantErr) && (!reflect.DeepEqual(Conf, tt.exCcfg)) {
				t.Errorf("Config does not match expected: %v, got %v", tt.exCcfg, Conf)
			}
			Conf.NhcConfig.Host = ""
			Conf.NhcConfig.Port = 0
			Conf.ServerConfig.LogLevel = ""
			Conf.JeedomConfig.URL = ""
			Conf.JeedomConfig.APIKey = ""
			Conf.JeedomConfig.Enabled = false
		})
	}
}
