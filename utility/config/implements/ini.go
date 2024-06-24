package implements

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/helper"
	"gopkg.in/ini.v1"
)

type iniConfig struct {
	handler *ini.File
	name    string
}

func (i *iniConfig) SetName(name_ string) {
	i.name = name_
}
func (i *iniConfig) Type() int {
	return constants.Ini
}

func (i *iniConfig) initialize() {
	var err error
	err = helper.CreateFileIfNotExist(i.GetConfigFileName())
	if err != nil {
		//log.Error(err)
		return
	}
	i.handler, err = ini.Load(i.GetConfigFileName())

	if err != nil {
		i.handler = ini.Empty()
	}
}

func (i *iniConfig) LoadToObject(section string, obj interface{}) error {
	//log.Info("Loading section:", section)

	if i.handler.HasSection(section) {
		if err := i.handler.Section(section).MapTo(obj); err != nil {
			//log.Errorf("%s faile to load. ERROR:%s", section, err)
			return err
		}
	} else {
		if err := i.handler.Section(section).ReflectFrom(obj); err != nil {
			//log.Errorf("%s fail to save. ERROR:%s", section, err)
			return err
		}
		i.Save()
	}
	return nil
}

func (i *iniConfig) UpdateToFile(section string, obj interface{}) error {
	i.handler.Section(section).ReflectFrom(obj)
	return i.Save()

}

func (i *iniConfig) Save() error {
	i.handler.SaveToIndent(i.GetConfigFileName(), "\t")
	return nil
}

func (i *iniConfig) GetConfigFileName() string {
	return constants.SettingFolder + "/" + i.name + constants.IniExt
}

func NewIniConfig() Interface.IConfiguration {
	i := &iniConfig{
		name: constants.SettingFileName,
	}
	i.initialize()
	return i
}
func NewIniConfigWithName(name_ string) Interface.IConfiguration {
	i := &iniConfig{
		name: name_,
	}
	i.initialize()
	return i
}
