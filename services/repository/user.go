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

	//Insert a user to the database
	// this method won't check if the user is exist or not
	InsertUser(user *models.User) error

	//Check if the user is exist
	// if the user is exist, return true
	// other wise return false
	HasUser(username string)

	//Check if the user and email is exist
	// if the user or email is exist, return true
	// other wise return false
	HasUserAndEmail(username, email string) bool
}

type user struct {
	db *gorp.DbMap
}

func (u *user) GetUser(id int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) GetUsers() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) GetUserByUsername(username string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) InsertUser(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *user) HasUser(username string) {
	//TODO implement me
	panic("implement me")
}

func (u *user) HasUserAndEmail(username, email string) bool {
	//TODO implement me
	panic("implement me")
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

func NewUserRepository(db *gorp.DbMap) UserRepository {
	u := &user{db}
	err := u.initTable()
	if err != nil {
		return nil
	}
	return u
}
