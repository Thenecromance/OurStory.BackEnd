package SQL

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Thenecromance/OurStories/base/SQL/Unit"
	"github.com/Thenecromance/OurStories/base/logger"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

var (
	inst *SQL
)

func init() {
	//if !runningOnOS() {
	//	start()
	//}
}

type SQL struct {
	*config

	defaultHandler *sql.DB
	handlerPool    map[string]*gorp.DbMap
}

func (s *SQL) getSafeHandler(db string) *gorp.DbMap {
	if db == "" {
		return nil
	}
	if conn, ok := s.handlerPool[db]; ok {
		return conn
	}

	logger.Get().Infof("could not find connection to [%s] prepare to create a new database", db)
	if err := s.createDatabase(db); err != nil {
		logger.Get().Errorf("could not create database [%s], Error:%s", db, err)
		return nil
	}

	s.handlerPool[db] = s.initGorpDb(db)

	s.DbName = append(s.DbName, db)
	return s.handlerPool[db]
}

func (s *SQL) init() {
	s.config = defaultConfig()
	s.defaultHandler = s.connectTo(s.DefaultDb)
	s.initAllConnection()
}
func (s *SQL) connectTo(db string) *sql.DB {
	handler, err := sql.Open(s.SqlType, s.connectStr(db))
	if err != nil {
		logger.Get().Errorf("failed to create handler to [%s] with error:%s", db, err)
		return nil
	}
	if err := handler.Ping(); err != nil {
		logger.Get().Errorf("failed to ping handler to [%s] with error:%s", db, err)
		return nil
	}
	return handler
}
func (s *SQL) initAllConnection() {
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
			logger.Get().Infof("skip default db [%s] could not create self!!!", dbName)
			continue
		}
		if s.createDatabase(dbName) == nil {
			s.handlerPool[dbName] = s.initGorpDb(dbName)
			logger.Get().Infof("[%s] handler coneected!", dbName)
		}
	}

}
func (s *SQL) createDatabase(db string) error {
	if s.defaultHandler == nil {
		return errors.New("default handler is nil")
	}
	logger.Get().Debugf("start to create database [%s] ...", db)
	script := "CREATE DATABASE  IF NOT EXISTS %s  CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
	unit := Unit.New(fmt.Sprintf(script, db))
	exec, err := s.defaultHandler.Exec(unit.Command())
	if err != nil {
		return err
	}
	id, _ := exec.LastInsertId()
	affected, _ := exec.RowsAffected()

	logger.Get().Debugf("[%s] create success! LastInsertId:%d, RowsAffected:%d", db, id, affected)
	return nil
}

// initialize a dbmap
func (s *SQL) initGorpDb(dbName string) *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish

	db, err := sql.Open(s.SqlType, s.connectStr(dbName))
	if err != nil {
		logger.Get().Errorf("sql.Open failed: %s", err)
	}
	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{
		Engine:   "InnoDB",
		Encoding: "utf8mb4",
	}}

	// dbmap.TraceOn("[gorp]", m.logger)
	return dbmap
}

func Instance() *SQL {
	if inst == nil {
		inst = &SQL{}
	}
	return inst
}

func Initialize() {
	Instance().init()
}

func Get(db string) *gorp.DbMap {
	return Instance().getSafeHandler(db)
}
func Default() *gorp.DbMap {
	return Instance().getSafeHandler("our_stories")
}
