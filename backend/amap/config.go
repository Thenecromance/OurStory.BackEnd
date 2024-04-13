package amap

import (
	"encoding/json"
	"os"
)

const (
	file = "./setting/aMap.json"
)

func (a *Amap) initConfig() {
	a.allowToUse = false
	if _, err := os.Stat(file); err != nil {

		log.Info("seems like you are first time to use this service")
		a.saveConfig()
		log.Infof("a new config file has been write to %s", file)
		return
	}

	log.Debugf("AMap found config file at %s start to load", file)
	a.loadConfig()

	a.allowToUse = true
}

func (a *Amap) loadConfig() {
	log.Debugf("start to load config...")
	buffer, err := os.ReadFile(file)
	if err != nil {
		log.Errorf("load config failed. with reason :%s", err.Error())
		a.allowToUse = true
		return
	}

	err = json.Unmarshal(buffer, a)
	if err != nil {
		log.Errorf("fail to unMarshal the config. with reason: %s", err.Error())
		a.allowToUse = false
		return
	}
	log.Debugf("config load complete!")
}

func (a *Amap) saveConfig() {
	log.Debugf("start to save the config....")

	buffer, err := json.Marshal(a)
	if err != nil {
		log.Errorf("failed to save the config... with reason: %s", err.Error())
		return
	}

	os.WriteFile(file, buffer, 0644)

	log.Debugf("save config complete!")
}
