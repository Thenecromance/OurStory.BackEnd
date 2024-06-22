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
	configPath = "setting/Config.ini"
)

var (
	defaultCfg *ini.File
)

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

func CloseIni() {
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

func MapTo(v interface{}) error {
	return defaultCfg.MapTo(v)
}
func MapSection(section string, v interface{}) error {
	err := defaultCfg.Section(section).MapTo(v)
	return err
}
func ReflectFrom(section string, v interface{}) (err error) {
	sec, err := defaultCfg.GetSection(section)
	sec.ReflectFrom(v)

	Flush()
	return
}

func HasSection(section string) bool {
	return defaultCfg.HasSection(section)
}
func GetString(section string, key string) string {
	return defaultCfg.Section(section).Key(key).String()
}
func GetStringWithDefault(section, key, shouldbe string) string {
	if defaultCfg.Section(section).HasKey(key) {
		return GetString(section, key)
	} else {
		SetString(section, key, shouldbe)
		return shouldbe
	}
}
func GetInt(section string, key string) int {
	return defaultCfg.Section(section).Key(key).MustInt()
}
func GetFloat32(section string, key string) float32 {
	return float32(GetFloat64(section, key))
}
func GetFloat64(section string, key string) float64 {
	return defaultCfg.Section(section).Key(key).MustFloat64()
}
func GetBool(section string, key string) bool {
	return defaultCfg.Section(section).Key(key).MustBool()
}
func GetUint(section string, key string) uint {
	return defaultCfg.Section(section).Key(key).MustUint()
}

func SetString(section string, key string, value string) {
	defaultCfg.Section(section).Key(key).SetValue(value)
}
func SetBool(section string, key string, value bool) {
	if value {
		defaultCfg.Section(section).Key(key).SetValue("true")
	} else {
		defaultCfg.Section(section).Key(key).SetValue("false")
	}
}

func Flush() {
	defaultCfg.SaveToIndent(configPath, "\t")
}
