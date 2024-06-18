package repository

import (
	"fmt"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type UserRepository interface {
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type user struct {
	db *gorp.DbMap
}

func (u *user) initTable() error {
	if u.db == nil {
		log.Debugf("db is nil")
		return fmt.Errorf("db is nil")
	}

	log.Infof("start to binding user with table user")
	tbl := u.db.AddTableWithName(models.User{}, "user")
	tbl.SetKeys(true, "Id") // using snowflake to generate the id

	err := u.db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table user with error: %s", err.Error())
		return err
	}

	return nil
}

/*
func NewUserRepository(db *gorp.DbMap) UserRepository {
	u := &user{db}
	u.initTable()
	return u
}
*/
