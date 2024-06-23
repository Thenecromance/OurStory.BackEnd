package Config

import (
	"errors"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/yaml.v3"
)

/*
a simply yaml config file implementation
this implementation is not thread safe, there is no lock mechanism
TODO: add lock mechanism
*/

type yamConfig struct {
	entries map[string]interface{} `yaml:"config"`
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
		log.Error(err)
		return
	}
	y.readContentFromFile()
}

func (y *yamConfig) readContentFromFile() {
	bytes, err := helper.ReadFile(y.GetConfigFileName())

	if err != nil {
		log.Error(err)
		return
	}

	if err = yaml.Unmarshal(bytes, y.entries); err != nil {
		log.Error(err)
		return
	}

	return
}

func (y *yamConfig) GetConfigFileName() string {
	return constants.SettingFolder + "/" + constants.SettingFileName + constants.YamlExt
}

func (y *yamConfig) LoadToObject(section string, obj interface{}) error {
	log.Debug("Loading section:", section)

	entry, exists := y.entries[section]
	if !exists {
		if err := y.UpdateToFile(section, obj); err != nil {
			log.Error("Failed to update file with section:", err)
			return err
		}
		return errors.New("section not found")
	}
	entryBytes, err := yaml.Marshal(entry)
	if err != nil {
		log.Error("Failed to marshal entry:", err)
		return err
	}

	if err := yaml.Unmarshal(entryBytes, obj); err != nil {
		log.Error("Failed to unmarshal entry to object:", err)
		return err
	}

	return nil
}

func (y *yamConfig) UpdateToFile(section string, obj interface{}) error {
	y.entries[section] = obj
	return y.Save()
}

func newYamlConfig() IConfiguration {
	y := &yamConfig{
		entries: make(map[string]interface{}),
	}
	y.initialize()
	return y
}
