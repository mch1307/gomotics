package config

import (
	"reflect"
	"testing"
	//. "github.com/mch1307/gomotics/config"
)

func TestInitialize(t *testing.T) {
	var cfg GlobalConfig
	cfg.NhcConfig.Host = "localhost"
	cfg.NhcConfig.Port = 8000
	cfg.ServerConfig.ListenPort = 8081
	cfg.ServerConfig.LogLevel = "DEBUG"
	cfg.ServerConfig.LogPath = "."

	tests := []struct {
		name     string
		confFile string
		wantErr  bool
		exCcfg   GlobalConfig
	}{
		{"read conf - listen port", "./test.toml", false, cfg},
		{"read conf - nhc host", "./test.toml", false, cfg},
		{"read conf - error", "./testt.toml", true, cfg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Initialize(tt.confFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			} else if (!tt.wantErr) && (!reflect.DeepEqual(Conf, tt.exCcfg)) {
				t.Errorf("Config does not match expect: %v, got %v", Conf, tt.exCcfg)
			}
		})
	}
}
