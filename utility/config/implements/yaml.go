package implements

import (
	"errors"
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/helper"
	"gopkg.in/yaml.v3"
)

/*
a simply yaml config file implementation
this implementation is not thread safe, there is no lock mechanism
TODO: add lock mechanism
*/

type yamConfig struct {
	entries map[string]interface{} `yaml:"config"`
	name    string
}

func (y *yamConfig) SetName(name string) {
	y.name = name
}

func (y *yamConfig) Type() int {
	return constants.Yaml
}

func (y *yamConfig) Save() error {
	bytes, err := yaml.Marshal(y.entries)
	if err != nil {
		return err
	}
	return helper.WriteFile(y.GetConfigFileName(), bytes)
}

func (y *yamConfig) initialize() {
	err := helper.CreateFileIfNotExist(y.GetConfigFileName())
	if err != nil {
		//log.Error(err)
		return
	}
	y.readContentFromFile()
}

func (y *yamConfig) readContentFromFile() {
	bytes, err := helper.ReadFile(y.GetConfigFileName())

	if err != nil {
		//log.Error(err)
		return
	}

	if err = yaml.Unmarshal(bytes, y.entries); err != nil {
		//log.Error(err)
		return
	}

	return
}

func (y *yamConfig) GetConfigFileName() string {
	return constants.SettingFolder + "/" + y.name + constants.YamlExt
}

func (y *yamConfig) LoadToObject(section string, obj interface{}) error {
	//log.Debug("Loading section:", section)

	entry, exists := y.entries[section]
	if !exists {
		if err := y.UpdateToFile(section, obj); err != nil {
			//log.Error("Failed to update file with section:", err)
			return err
		}
		return errors.New("section not found")
	}
	entryBytes, err := yaml.Marshal(entry)
	if err != nil {
		//log.Error("Failed to marshal entry:", err)
		return err
	}

	if err = yaml.Unmarshal(entryBytes, obj); err != nil {
		//log.Error("Failed to unmarshal entry to object:", err)
		return err
	}

	return nil
}

func (y *yamConfig) UpdateToFile(section string, obj interface{}) error {
	y.entries[section] = obj
	return y.Save()
}

func NewYamlConfig() Interface.IConfiguration {
	y := &yamConfig{
		entries: make(map[string]interface{}),
		name:    constants.SettingFileName,
	}
	y.initialize()
	return y
}

func NewYamlConfigWithName(name string) Interface.IConfiguration {
	y := &yamConfig{
		entries: make(map[string]interface{}),
		name:    name,
	}
	y.initialize()
	return y
}
