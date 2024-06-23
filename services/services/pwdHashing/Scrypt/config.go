package Scrypt

import (
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
)

const (
	sectionName = "Hashing.scrypt"
)

// Setting is the configuration for the scrypt algorithm
type Setting struct {
	CostFactor            int `ini:"cost_factor" json:"cost_factor,omitempty" yaml:"cost_factor"`
	BlockSizeFactor       int `ini:"block_size" json:"block_size_factor,omitempty" yaml:"block_size_factor"`
	ParallelizationFactor int `ini:"parallelization_factor" json:"parallelization_factor,omitempty" yaml:"parallelization_factor"`
	KeyLen                int `ini:"key_len" json:"key_len,omitempty" yaml:"key_len"`
	RandomSaltLen         int `ini:"random_salt_len" json:"random_salt_len,omitempty" yaml:"random_salt_len"`
}

func defaultConfig() *Setting {
	return &Setting{
		CostFactor:            16384,
		BlockSizeFactor:       8,
		ParallelizationFactor: 1,
		KeyLen:                32,
		RandomSaltLen:         16,
	}
}

func newConfig() *Setting {
	var cfg *Setting
	cfg = defaultConfig()

	err := Config.Instance(Config.Yaml).LoadToObject(sectionName, cfg)
	if err != nil {
		log.Warnf("%s, start to use default instead", err)
		cfg = defaultConfig()
	}

	return cfg
}
