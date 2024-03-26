package AMapToken

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
)

var instance *Token

func init() {

}

type Token struct {
	Amap string `ini:"amap" json:"amap"`
}

func (t *Token) Load() {
	if err := Config.MapSection("Tokens", t); err != nil {
		logger.Get().Errorf("%s faile to map section. ERROR:%s", "Tokens", err)
		return
	}
	if err := Config.ReflectFrom("Tokens", t); err != nil {
		logger.Get().Errorf("%s faile to reflect section. ERROR:%s", "Tokens", err)
		return
	}
}

func Instance() *Token {
	if instance == nil {
		instance = &Token{}
		instance.Load()
	}
	return instance
}
