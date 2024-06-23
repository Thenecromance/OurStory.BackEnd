package Config

import (
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/ini.v1"
)

type iniConfig struct {
	handler *ini.File
}

func (i *iniConfig) initialize() {
	var err error
	err = helper.CreateFileIfNotExist(i.GetConfigFileName())
	if err != nil {
		log.Error(err)
		return
	}
	i.handler, err = ini.Load(i.GetConfigFileName())

	if err != nil {
		i.handler = ini.Empty()
	}
}

func (i *iniConfig) LoadToObject(section string, obj interface{}) error {
	log.Info("Loading section:", section)

	/*	//try to load the default config
		if err := i.handler.Section(section).MapTo(obj); err != nil {
			log.Errorf("%s faile to load. ERROR:%s", section, err)
			// save the default config to file
			if err := i.handler.Section(section).ReflectFrom(obj); err != nil {
				log.Errorf("%s fail to save. ERROR:%s", section, err)
				return errors.New("failed to save section")
			}
			return i.Save()
		}
		return nil*/
	if i.handler.HasSection(section) {
		if err := i.handler.Section(section).MapTo(obj); err != nil {
			log.Errorf("%s faile to load. ERROR:%s", section, err)
			return err
		}
	} else {
		if err := i.handler.Section(section).ReflectFrom(obj); err != nil {
			log.Errorf("%s fail to save. ERROR:%s", section, err)
			return err
		}
		i.Save()
	}
	return nil
}

func (i *iniConfig) UpdateToFile(section string, obj interface{}) error {
	i.handler.Section(section).ReflectFrom(obj)
	return i.Save()
	/*if i.handler.HasSection(section) {
		i.handler.Section(section).ReflectFrom(obj)
		i.Save()
		return nil
	}
	return nil*/
}

func (i *iniConfig) Save() error {
	i.handler.SaveToIndent(i.GetConfigFileName(), "\t")
	return nil
}

func (i *iniConfig) GetConfigFileName() string {
	return constants.SettingFolder + "/" + constants.SettingFileName + constants.IniExt
}

func newIniConfig() IConfiguration {
	i := &iniConfig{}
	i.initialize()
	return i
}
