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
	cfgDefault.NhcConfig.Host = ""
	cfgDefault.NhcConfig.Port = 0
	cfgDefault.ServerConfig.ListenPort = 8081
	cfgDefault.ServerConfig.LogLevel = "INFO"
	cfgDefault.ServerConfig.LogPath, _ = os.Getwd()
	tests := []struct {
		name     string
		confFile string
		wantErr  bool
		exCcfg   GlobalConfig
	}{
		{"read conf - listen port", "test.toml", false, cfg},
		{"read conf - nhc host", "test.toml", false, cfg},
		{"read conf - error", "NoSuchFile", false, cfgDefault},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Initialize(tt.confFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			} else if (!tt.wantErr) && (!reflect.DeepEqual(Conf, tt.exCcfg)) {
				t.Errorf("Config does not match expected: %v, got %v", tt.exCcfg, Conf)
			}
			Conf.NhcConfig.Host = ""
			Conf.NhcConfig.Port = 0
		})
	}
}
