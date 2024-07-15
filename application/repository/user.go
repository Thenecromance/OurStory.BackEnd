package repository

import (
	"fmt"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type UserRepository interface {
	Interface.Repository
	GetUser(id int) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByUsername(username string) (*models.User, error)

	GetUserIdByName(username string) (int64, error)
	// InsertUser Insert a user to the database
	// this method won't check if the user is exist or not
	InsertUser(user *models.User) error

	// HasUser Check if the user is exist
	// if the user is exist, return true
	// other wise return false
	HasUser(username string)

	HasId(id int) bool

	// HasUserAndEmail Check if the user and email is exist
	// if the user or email is exist, return true
	// other wise return false
	HasUserAndEmail(username, email string) bool
	UpdateLastLogin(id int64, unix int64) error
}

type user struct {
	db *gorp.DbMap
}

func (u *user) BindTable() error {
	u.db.AddTableWithName(models.User{}, "Users")
	return nil
}

func (u *user) GetUserIdByName(username string) (int64, error) {

	selectInt, err := u.db.SelectInt("select id from user where username = ?", username)
	if err != nil {
		return -1, err
	}

	return selectInt, nil
}

func (u *user) GetUser(id int) (*models.User, error) {
	//models.User
	obj, err := u.db.Select(models.User{}, "select * from user where id = ?", id)
	if err != nil {
		return nil, err
	}
	return obj[0].(*models.User), nil
}

func (u *user) GetUsers() ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *user) GetUserByUsername(username string) (*models.User, error) {

	obj, err := u.db.Select(models.User{}, "select * from user where username = ?", username)
	if err != nil {
		return nil, err
	}
	if len(obj) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return obj[0].(*models.User), nil
}

func (u *user) UpdateLastLogin(id int64, unix int64) error {
	trans, err := u.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction with error: %s", err.Error())
		return err
	}

	_, err = trans.Exec("update user set last_login = ? where id = ?", unix, id)
	if err != nil {
		log.Errorf("failed to update last login with error: %s", err.Error())
		trans.Rollback()
		return err
	}

	return trans.Commit()
}

func (u *user) InsertUser(user *models.User) error {
	transaction, err := u.db.Begin()
	if err != nil {
		return err
	}
	err = transaction.Insert(user)
	if err != nil {
		transaction.Rollback()
		log.Errorf("failed to insert user with error: %s", err.Error())
		return err
	}

	return transaction.Commit()
}

func (u *user) HasUser(username string) {
	obj, err := u.db.SelectInt("select count(*) from user where username = ?", username)
	if err != nil {
		return
	}
	if obj > 0 {
		return

	}
}

func (u *user) HasId(id int) bool {
	obj, err := u.db.SelectInt("select count(*) from user where id = ?", id)
	if err != nil {
		return false
	}
	if obj > 0 {
		return true
	}
	return false
}

func (u *user) HasUserAndEmail(username, email string) bool {
	obj, err := u.db.SelectInt("select count(*) from user where username = ? or email = ?", username, email)
	if err != nil {
		return false
	}
	if obj > 0 {
		return true
	}
	return false
}

/*func (u *user) initTable() error {
	if u.db == nil {
		log.Debugf("db is nil")
		return fmt.Errorf("db is nil")
	}

	log.Infof("start to binding user with table user")
	tbl := u.db.AddTableWithName(models.User{}, "Users")
	tbl.SetKeys(true, "UserId") // using snowflake to generate the id

	err := u.db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table user with error: %s", err.Error())
		return err
	}
	return nil
}*/

func NewUserRepository(db *gorp.DbMap) UserRepository {
	u := &user{db}
	/*	err := u.initTable()
		if err != nil {
			return nil
		}*/
	return u
}
