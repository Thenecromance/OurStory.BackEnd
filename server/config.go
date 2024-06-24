package server

import (
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
)

type config struct {
	Addr                         string `ini:"addr" json:"addr" yaml:"addr"`
	ReadTimeout                  int    `ini:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout                 int    `ini:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
	IdleTimeout                  int    `ini:"idle_timeout" json:"idle_timeout" yaml:"idle_timeout"`
	MaxHeaderBytes               int    `ini:"max_header_bytes" json:"max_header_bytes" yaml:"max_header_bytes"`
	DisableGeneralOptionsHandler bool   `ini:"disable_general_options_handler" json:"disable_general_options_handler" yaml:"disable_general_options_handler"`
}

func (cfg *config) defaultConfig() {
	cfg.Addr = ":8080"
	cfg.ReadTimeout = 10
	cfg.WriteTimeout = 10
	cfg.IdleTimeout = 10
	cfg.MaxHeaderBytes = 1 << 20
	cfg.DisableGeneralOptionsHandler = false
}

func (cfg *config) load() {
	cfg.defaultConfig()
	err := /*Config.LoadToObject("server", cfg)*/ Config.Instance().LoadToObject("server", cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
