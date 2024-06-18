package SQL

import (
	"fmt"
	Config "github.com/Thenecromance/OurStories/utility/config"
	"github.com/Thenecromance/OurStories/utility/log"
)

type config struct {
	SqlType   string   `ini:"sql" comment:"which sql to use, default mysql"`
	Protocol  string   `ini:"protocol"  comment:"tcp or unix socket"`
	Host      string   `ini:"host" `
	Port      string   `ini:"port"`
	User      string   `ini:"user"`
	Password  string   `ini:"password"`
	DefaultDb string   `ini:"default_db"  comment:"when there is no db_name, use this,and it need to create first"`
	DbName    []string `ini:"db_name" comment:"the logic database name which should be the same as the database's name'"`
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
		DbName:    []string{},
	}
	err := Config.LoadToObject(sectionName, cfg)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	log.Infof("LoadController config section to %s", sectionName)
	return cfg
}
