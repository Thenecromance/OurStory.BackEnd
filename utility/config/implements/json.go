package implements

import (
	"encoding/json"
	"errors"
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/helper"
)

type jsonConfig struct {
	entries map[string]interface{}
	name    string
}

func (j *jsonConfig) SetName(name string) {
	j.name = name
}

func (j *jsonConfig) Type() int {
	return constants.Json
}

func (j *jsonConfig) initialize() {
	err := helper.CreateFileIfNotExist(j.GetConfigFileName())
	if err != nil {
		//log.Error(err)
		return
	}
	j.readContentFromFile()
}

func (j *jsonConfig) readContentFromFile() {
	bytes, err := helper.ReadFile(j.GetConfigFileName())

	if err != nil {
		//log.Error(err)
		return
	}

	if err = json.Unmarshal(bytes, &j.entries); err != nil {
		//log.Error(err)
		return
	}

	return
}

func (j *jsonConfig) LoadToObject(section string, obj interface{}) error {
	//log.Debug("Loading section:", section)

	entry, exists := j.entries[section]
	if !exists {
		if err := j.UpdateToFile(section, obj); err != nil {
			//log.Error("Failed to update file with section:", err)
			return err
		}
		return errors.New("section not found")
	}
	entryBytes, err := json.Marshal(entry)
	if err != nil {
		//log.Error("Failed to marshal entry:", err)
		return err
	}

	if err := json.Unmarshal(entryBytes, obj); err != nil {
		//log.Error("Failed to unmarshal entry to object:", err)
		return err
	}

	return nil
}

func (j *jsonConfig) UpdateToFile(section string, obj interface{}) error {
	j.entries[section] = obj
	return j.Save()
}

func (j *jsonConfig) GetConfigFileName() string {
	return constants.SettingFolder + "/" + j.name + constants.JsonExt
}

func (j *jsonConfig) Save() error {
	bytes, err := json.MarshalIndent(j.entries, "", "  ")
	if err != nil {
		return err
	}
	return helper.WriteFile(j.GetConfigFileName(), bytes)
}

func NewJsonConfig() Interface.IConfiguration {
	j := &jsonConfig{
		entries: make(map[string]interface{}),
		name:    constants.SettingFileName,
	}
	j.initialize()
	return j
}

func NewJsonConfigWithName(name string) Interface.IConfiguration {
	j := &jsonConfig{
		entries: make(map[string]interface{}),
		name:    name,
	}
	j.initialize()
	return j
}
