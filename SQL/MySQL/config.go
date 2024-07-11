package MySQL

import (
	"fmt"
	Config "github.com/Thenecromance/OurStories/utility/config"

	"github.com/Thenecromance/OurStories/utility/log"
)

type config struct {
	SqlType   string `ini:"sql"           json:"sql_type"   yaml:"sql_type"    `
	Protocol  string `ini:"protocol"      json:"protocol"   yaml:"protocol"    `
	Host      string `ini:"host"          json:"host"       yaml:"host"`
	Port      string `ini:"port"          json:"port"       yaml:"port"`
	User      string `ini:"user"          json:"user"       yaml:"user"`
	Password  string `ini:"password"      json:"password"   yaml:"password"`
	DefaultDb string `ini:"default_db"    json:"default_db" yaml:"default_db"  `
}

func (c *config) buildConnectString(dbName string) (res string) {

	res = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4",
		c.User,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		dbName,
	)
	log.Debugf("connect string is %s", res)
	return
}

// defaultConfig is a function to create a default config
func (c *config) buildDefaultConnectString() string {
	return c.buildConnectString(c.DefaultDb)
}

func defaultConfig() *config {
	const (
		sectionName = "SQLManager"
	)
	//default config
	cfg := &config{
		SqlType:   "mysql",
		Protocol:  "tcp",
		Host:      "127.0.0.1",
		Port:      "3306",
		User:      "root",
		DefaultDb: "mysql",
	}
	err := Config.Instance().LoadToObject(sectionName, &cfg)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	log.Infof("LoadController config section to %s", sectionName)
	return cfg
}
