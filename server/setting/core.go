package setting

import (
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
	"time"
)

// Core Setting
type Core struct {
	Addr           string        `json:"addr"               yaml:"addr"`
	ReadTimeout    time.Duration `json:"read_timeout"       yaml:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"      yaml:"write_timeout"`
	IdleTimeout    time.Duration `json:"idle_timeout"       yaml:"idle_timeout"`
	MaxHeaderBytes int           `json:"max_header_bytes"   yaml:"max_header_bytes"`
	// TLS certificate and key file path if not provided, server will run in http mode
	KeyPath  string `json:"key_path" yaml:"key_path"`
	CertPath string `json:"cert_path" yaml:"cert_path"`
}

func (cfg *Core) defaultConfig() {
	cfg.Addr = ":8080"
	cfg.ReadTimeout = 10 * time.Second
	cfg.WriteTimeout = 10 * time.Second
	cfg.IdleTimeout = 10 * time.Second
	cfg.MaxHeaderBytes = 1 << 20
	//cfg.DisableGeneralOptionsHandler = false
}

func (cfg *Core) Load() {
	cfg.defaultConfig()
	err := /*Config.LoadToObject("server", cfg)*/ Config.Instance().LoadToObject("server", cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
