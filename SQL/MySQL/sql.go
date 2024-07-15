package MySQL

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Thenecromance/OurStories/utility/File"
	"github.com/Thenecromance/OurStories/utility/log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

var (
	inst *SQLManager
)

func init() {
	//if !runningOnOS() {
	//	start()
	//}
}

type SQLManager struct {
	*config
	DbName         []string
	defaultHandler *sql.DB
	handlerPool    map[string]*gorp.DbMap
}

func (s *SQLManager) getSafeHandler(db string) *gorp.DbMap {
	if db == "" {
		db = DEFATUL_DB_NAME
	}
	if conn, ok := s.handlerPool[db]; ok {
		return conn
	}

	log.Infof("could not find connection to [%s] prepare to connect to database", db)
	if err := s.createDatabase(db); err != nil {
		log.Errorf("could not create database [%s], Error:%s", db, err)
		return nil
	}

	s.handlerPool[db] = s.initGorpDb(db)

	s.DbName = append(s.DbName, db)
	return s.handlerPool[db]
}

func (s *SQLManager) init() {
	s.config = defaultConfig()
	s.defaultHandler = s.connectTo(s.DefaultDb)
	s.initAllConnection()
}
func (s *SQLManager) connectTo(db string) *sql.DB {
	handler, err := sql.Open(s.SqlType, s.buildConnectString(db))
	if err != nil {
		log.Errorf("failed to create handler to [%s] with error:%s", db, err)
		return nil
	}
	if err := handler.Ping(); err != nil {
		log.Errorf("failed to ping handler to [%s] with error:%s", db, err)
		return nil
	}
	return handler
}

// initialize all connection which stored in config files
func (s *SQLManager) initAllConnection() {
	if s.handlerPool == nil {
		s.handlerPool = make(map[string]*gorp.DbMap)
	}

	if len(s.DbName) == 0 {
		return
	}

	for _, dbName := range s.DbName {
		if dbName == "" {
			continue
		}
		if dbName == s.DefaultDb {
			log.Infof("skip default db [%s] could not create self!!!", dbName)
			continue
		}
		if s.createDatabase(dbName) == nil {
			s.handlerPool[dbName] = s.initGorpDb(dbName)
			log.Infof("[%s] handler coneected!", dbName)
		}
	}

}

// when database is not exists , just create it by using default handler
func (s *SQLManager) createDatabase(db string) error {
	if s.defaultHandler == nil {
		return errors.New("default handler is nil")
	}
	log.Debugf("[%s] start to create database ...", db)
	script := "CREATE DATABASE  IF NOT EXISTS %s  CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"

	exec, err := s.defaultHandler.Exec(fmt.Sprintf(script, db))
	if err != nil {
		return err
	}
	id, _ := exec.LastInsertId()
	affected, _ := exec.RowsAffected()

	log.Debugf("[%s] create success! LastInsertId:%d, RowsAffected:%d", db, id, affected)
	return nil
}

// initialize a dbmap
func (s *SQLManager) initGorpDb(dbName string) *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish

	db, err := sql.Open(s.SqlType, s.buildConnectString(dbName))
	if err != nil {
		log.Errorf("sql.Open failed: %s", err)
	}
	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{
		Engine:   "InnoDB",
		Encoding: "utf8mb4",
	}}

	dbmap.TraceOn("[MySQL Query]", log.Instance)
	return dbmap
}

// Instance  using lazy mode to get sql connections' manager
func Instance() *SQLManager {
	if inst == nil {
		inst = &SQLManager{}
		inst.init()
	}
	return inst
}

// Get a connection to a specific database
func Get(db string) *gorp.DbMap {
	return Instance().getSafeHandler(db)
}

// Default will return the default database connection
func Default() *gorp.DbMap {
	return Instance().getSafeHandler(DEFATUL_DB_NAME)
}

func RunScript(script string) error {
	return RunScriptOnDb(DEFATUL_DB_NAME, script)
}

func RunScriptOnDb(db, script string) error {
	buffer, _ := File.ReadFrom(script)
	_, err := Get(db).Exec(string(buffer))
	return err
}
