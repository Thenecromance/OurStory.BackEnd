package Config

/*
Interface segregation principle
*/

/*
IConfiguration is an interface that defines the methods that a configuration should implement
*/
type IConfiguration interface {
	// LoadToObject loads the configuration file into the object
	//
	// @param: section: the section name of the config
	//
	// @param: obj: the object to load the config to
	LoadToObject(section string, obj interface{}) error

	// UpdateToFile updates the configuration file with the new values
	//
	// @param: section: the section name of the config
	//
	// @param: obj: the object to update the config from
	UpdateToFile(section string, obj interface{}) error

	//Save saves the configuration file
	Save() error

	// GetConfigFileName returns the name of the configuration file
	GetConfigFileName() string
}

const (
	Ini = iota
	Json
	Yaml
)

var (
	defaultInst IConfiguration

	//different Instances to support different config types
	iniInst  IConfiguration
	jsonInst IConfiguration
	yamlInst IConfiguration
)

func New(configType int) IConfiguration {
	switch configType {
	case Ini:
		return newIniConfig()
	case Json:
		return newJsonConfig()
	case Yaml:
		return newYamlConfig()
	default:
		panic("Invalid config type")
	}
}

func Instance() IConfiguration {
	if defaultInst == nil {
		SetDefault(Yaml)
	}
	return defaultInst
}

func SetDefault(configType int) {
	switch configType {
	case Ini:
		if iniInst == nil {
			iniInst = newIniConfig()
		}
		defaultInst = iniInst
	case Json:
		if jsonInst == nil {
			jsonInst = newJsonConfig()
		}
		defaultInst = jsonInst
	case Yaml:
		if yamlInst == nil {
			yamlInst = newYamlConfig()
		}
		defaultInst = yamlInst
	default:
		panic("Invalid config type")
	}
}
