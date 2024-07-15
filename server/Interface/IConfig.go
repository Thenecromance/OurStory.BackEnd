package Interface

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

	SetName(name string)

	Type() int
}
