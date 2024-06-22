package Scrypt

import (
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
)

const (
	sectionName = "scrypt"
)

// Setting is the configuration for the scrypt algorithm
type Setting struct {
	N             int `ini:"n"`
	R             int `ini:"r"`
	P             int `ini:"p"`
	KeyLen        int `ini:"key_len"`
	RandomSaltLen int `ini:"random_salt_len"`
}

func defaultConfig() *Setting {
	return &Setting{
		N:             16384,
		R:             8,
		P:             1,
		KeyLen:        32,
		RandomSaltLen: 16,
	}
}

func newConfig() *Setting {
	var cfg *Setting
	cfg = defaultConfig()
	if !Config.HasSection(sectionName) {
		Config.MapSection(sectionName, cfg)
		Config.Flush()
		return cfg
	}

	err := Config.LoadToObject("scrypt", cfg)
	if err != nil {
		log.Error(err)
		cfg = defaultConfig()
	}

	return cfg
}
