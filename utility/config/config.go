package Config

import (
	"errors"
	"fmt"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"

	"gopkg.in/ini.v1"
	"reflect"
	"unsafe"
)

const (
	configPath = "setting/Config.ini" //the path of the config file
)

var (
	defaultCfg *ini.File // the default config all the config operation will be based on this
)

// init will load the default config from the file
func init() {
	var err error
	err = helper.CreateIfNotExist("./setting")
	if err != nil {
		log.Error(err)
		return
	}
	defaultCfg, err = ini.Load(configPath)

	if err != nil {
		defaultCfg = ini.Empty()
	}
}

// CloseIni will save the config to the file
// just need to be called when the program is about to exit
func CloseIni() {
	defaultCfg.SaveToIndent(configPath, "\t")
}

// Flush will save the raw config to the file
func Flush() {
	defaultCfg.SaveToIndent(configPath, "\t")
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// LoadToObject is a method to directly load the config to the object
//
// @param: section: the section name of the config
//
// @param: obj: the object to load the config to
//
// @return: error when the section is not found
func LoadToObject(section string, obj interface{}) error {
	//try to load the default config
	if err := MapSection(section, obj); err != nil {
		log.Errorf("%s faile to load. ERROR:%s", section, err)
		return errors.New("ini has no section")
	}
	// save the default config to file
	if err := ReflectFrom(section, obj); err != nil {
		log.Errorf("%s faile to save. ERROR:%s", section, err)
		return errors.New("faile to save section")
	}
	return nil
}

func AppendSection(section string, format string, args ...interface{}) error {
	if HasSection(section) {
		return nil
	}
	format = fmt.Sprintf(format, args...)
	log.Debug("AppendSection: %s", format)
	return defaultCfg.Append(String2Bytes(format))
}

// MapTo will map the config to the object
//
// @param: v: the object to map the config to
//
// @return: error when the object is not a pointer
func MapTo(v interface{}) error {
	return defaultCfg.MapTo(v)
}

// MapSection will map the section of the config to the object
//
// @param: section: the section name of the config
//
// @param: v: the object to map the config to
func MapSection(section string, v interface{}) error {
	err := defaultCfg.Section(section).MapTo(v)
	return err
}

// ReflectFrom will reflect the object to the config
//
// @param: section: the section name of the config
//
// @param: v: the object to reflect the config from
func ReflectFrom(section string, v interface{}) (err error) {
	sec, err := defaultCfg.GetSection(section)
	sec.ReflectFrom(v)

	Flush()
	return
}

// HasSection will check if config has the section
//
// @param: section: the section name of the config
//
// @return: true if the section is found, false otherwise
func HasSection(section string) bool {
	return defaultCfg.HasSection(section)
}

// GetString will get the section of the config
//
// @param: section: the section name of the config
//
// @return: the section of the config
func GetString(section string, key string) string {
	return defaultCfg.Section(section).Key(key).String()
}

// GetStringWithDefault will get the section of the config
// if the section is not found, it will set the section to the default value
//
// @param: section: the section name of the config
//
// @param: key: the key of the config
func GetStringWithDefault(section, key, shouldbe string) string {
	if defaultCfg.Section(section).HasKey(key) {
		return GetString(section, key)
	} else {
		SetString(section, key, shouldbe)
		return shouldbe
	}
}

// GetInt will get the section of the config
//
// @param: section: the section name of the config
//
// @return: the section of the config
//
// @return: the section of the config
func GetInt(section string, key string) int {
	return defaultCfg.Section(section).Key(key).MustInt()
}

// GetFloat32 will get the section of the config
//
// @param: section: the section name of the config
//
// @param: key: the key of the config
func GetFloat32(section string, key string) float32 {
	return float32(GetFloat64(section, key))
}

// GetFloat64 will get the section of the config
//
// @param: section: the section name of the config
//
// @param: key: the key of the config
func GetFloat64(section string, key string) float64 {
	return defaultCfg.Section(section).Key(key).MustFloat64()
}

// GetBool will get the section of the config
//
// @param: section: the section name of the config
//
// @param: key: the key of the config
func GetBool(section string, key string) bool {
	return defaultCfg.Section(section).Key(key).MustBool()
}

// GetUint will get the section of the config
//
// @param: section: the section name of the config
//
// @param: key: the key of the config
func GetUint(section string, key string) uint {
	return defaultCfg.Section(section).Key(key).MustUint()
}

/*
SetString will set the section of the config

@param: section: the section name of the config

@param: key: the key of the config

@param: value: the value of the config
*/
func SetString(section string, key string, value string) {
	defaultCfg.Section(section).Key(key).SetValue(value)
}

/*
SetBool will set the section of the config
*/
func SetBool(section string, key string, value bool) {
	if value {
		defaultCfg.Section(section).Key(key).SetValue("true")
	} else {
		defaultCfg.Section(section).Key(key).SetValue("false")
	}
}
