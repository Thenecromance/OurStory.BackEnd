package Argon2

import (
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
)

const (
	sectionName = "Hashing.argon2"
)

type Setting struct {
	Threads uint32 `ini:"threads"`
	Memory  uint32 `ini:"memory"`
	Time    uint32 `ini:"time"`
	KeyLen  uint32 `ini:"key_len"`

	RandomSaltLen int `ini:"random_salt_len"`
}

func defaultSetting() *Setting {
	return &Setting{
		Threads: 4,
		Memory:  64 * 1024,
		Time:    1,
		KeyLen:  32,

		RandomSaltLen: 16,
	}
}
func newSetting() *Setting {
	var cfg *Setting
	cfg = defaultSetting()
	/*	if !Config.HasSection(sectionName) {
		Config.MapSection(sectionName, cfg)
		Config.Flush()
		return cfg
	}*/

	err := Config.LoadToObject(sectionName, cfg)
	if err != nil {
		log.Error(err)
		cfg = defaultSetting()
	}

	return cfg
}
