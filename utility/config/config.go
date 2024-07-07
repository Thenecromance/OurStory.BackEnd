package Config

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/utility/config/implements"
)

var (
	defaultInst Interface.IConfiguration

	//different Instances to support different config types
	jsonInst Interface.IConfiguration
	yamlInst Interface.IConfiguration

	customConfig map[string]Interface.IConfiguration
)

func New(configType int) Interface.IConfiguration {
	switch configType {
	/*case constants.Ini:
	return implements.NewIniConfig()*/
	case constants.Json:
		return implements.NewJsonConfig()
	case constants.Yaml:
		return implements.NewYamlConfig()
	default:
		panic("Invalid config type")
	}
}
func NewWithName(name string, configType int) Interface.IConfiguration {
	switch configType {
	/*case constants.Ini:
	return implements.NewIniConfigWithName(name)*/
	case constants.Json:
		return implements.NewJsonConfigWithName(name)
	case constants.Yaml:
		return implements.NewYamlConfigWithName(name)
	default:
		panic("Invalid config type")
	}

}

func Instance() Interface.IConfiguration {
	if defaultInst == nil {
		SetDefault(constants.Yaml)
	}
	return defaultInst
}

func SetDefault(configType int) {
	switch configType {
	/*case constants.Ini:
	if iniInst == nil {
		iniInst = implements.NewIniConfig()
	}
	defaultInst = iniInst*/
	case constants.Json:
		if jsonInst == nil {
			jsonInst = implements.NewJsonConfig()
		}
		defaultInst = jsonInst
	case constants.Yaml:
		if yamlInst == nil {
			yamlInst = implements.NewYamlConfig()
		}
		defaultInst = yamlInst
	default:
		panic("Invalid config type")
	}
}

func InstanceByName(name string, type_ int) Interface.IConfiguration {
	if customConfig == nil {
		customConfig = make(map[string]Interface.IConfiguration)
	}
	if customConfig[name] == nil {
		//customConfig[name] = implements.NewJsonConfig()
		customConfig[name] = NewWithName(name, type_)
	}
	return customConfig[name]
}

func LoadToObject(section string, obj interface{}) error {
	return Instance().LoadToObject(section, obj)
}
