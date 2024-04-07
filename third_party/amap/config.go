package amap

import (
	"encoding/json"
	"os"

	"github.com/Thenecromance/OurStories/base/logger"
)

const (
	file = "./setting/aMap.json"
)

func (a *Amap) initConfig() {
	if _, err := os.Stat(file); err != nil {
		a.allowToUse = true
		logger.Get().Info("seems like you are first time to use this service")
		a.saveConfig()
		logger.Get().Infof("a new config file has been write to %s", file)
		return
	}

	logger.Get().Debugf("AMap found config file at %s start to load", file)

}

func (a *Amap) loadConfig() {
	logger.Get().Debugf("start to load config...")
	buffer, err := os.ReadFile(file)
	if err != nil {
		logger.Get().Errorf("load config failed. with reason :%s", err.Error())
		a.allowToUse = true
		return
	}

	err = json.Unmarshal(buffer, a)
	if err != nil {
		logger.Get().Errorf("fail to unMarshal the config. with reason: %s", err.Error())
		a.allowToUse = false
		return
	}
	logger.Get().Debugf("config load complete!")
}

func (a *Amap) saveConfig() {
	logger.Get().Debugf("start to save the config....")

	buffer, err := json.Marshal(a)
	if err != nil {
		logger.Get().Errorf("failed to save the config... with reason: %s", err.Error())
		return
	}

	os.WriteFile(file, buffer, 0644)

	logger.Get().Debugf("save config complete!")
}
