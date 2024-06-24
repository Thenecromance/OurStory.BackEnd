package constants

const (
	Ini = iota
	Json
	Yaml
)

const (
	SettingFolder   = "settings"
	SettingFileName = "Config" // don't contain the file extension, just let the user specify the file extension

	YamlExt = ".yaml"
	JsonExt = ".json"
	IniExt  = ".ini"
)
