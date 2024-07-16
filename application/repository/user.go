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
	GetUser(id int64) (*models.User, error)
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

	HasId(id int64) bool

	// HasUserAndEmail Check if the user and email is exist
	// if the user or email is exist, return true
	// other wise return false
	HasUserAndEmail(username, email string) bool
	UpdateLastLogin(id int64, unix int64) error
}

const (
	UserDbName = "Users"
)

type user struct {
	db *gorp.DbMap
}

func (u *user) BindTable() error {
	u.db.AddTableWithName(models.User{}, "Users").SetKeys(true, "UserId")
	u.db.AddTableWithName(models.LoginLogs{}, "LoginLogs").SetKeys(false, "UserId")
	return nil
}

func (u *user) GetUserIdByName(username string) (int64, error) {

	selectInt, err := u.db.SelectInt("select user_id from "+UserDbName+" where username = ?", username)
	if err != nil {
		return -1, err
	}

	return selectInt, nil
}

func (u *user) GetUser(id int64) (*models.User, error) {
	//models.User
	obj, err := u.db.Select(models.User{}, "select * from Users where user_id = ?", id)
	if err != nil {
		return nil, err
	}
	return obj[0].(*models.User), nil
}

// do not use this by client this will return all data in the table
func (u *user) GetUsers() ([]models.User, error) {
	var users []models.User
	_, err := u.db.Select(&users, "select * from Users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *user) GetUserByUsername(username string) (*models.User, error) {

	obj, err := u.db.Select(models.User{}, "select * from Users where username = ?", username)
	if err != nil {
		return nil, err
	}
	if len(obj) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return obj[0].(*models.User), nil
}

func (u *user) UpdateLastLogin(id int64, unix int64) error {
	//trans, err := u.db.Begin()
	//if err != nil {
	//	log.Errorf("failed to begin transaction with error: %s", err.Error())
	//	return err
	//}
	//
	///*_, err = trans.Exec("update Users set last_login = ? where user_id = ?", unix, id)*/
	//_, err = trans.Update(&models.LoginLogs{UserId: id})
	//if err != nil {
	//	log.Errorf("failed to update last login with error: %s", err.Error())
	//	trans.Rollback()
	//	return err
	//}
	//
	//return trans.Commit()

	log.Info("update last login", id)
	_, err := u.db.Update(&models.LoginLogs{UserId: id})
	if err != nil {
		return err
	}
	return nil
}

func (u *user) InsertUser(user *models.User) error {

	err := u.db.Insert(user)
	if err != nil {

		log.Errorf("failed to insert user with error: %s", err.Error())
		return err
	}
	return nil
}

func (u *user) HasUser(username string) {
	obj, err := u.db.SelectInt("select count(*) from Users where username = ?", username)
	if err != nil {
		return
	}
	if obj > 0 {
		return

	}
}

func (u *user) HasId(id int64) bool {
	obj, err := u.db.SelectInt("select count(*) from Users where id = ?", id)
	if err != nil {
		return false
	}
	if obj > 0 {
		return true
	}
	return false
}

func (u *user) HasUserAndEmail(username, email string) bool {
	obj, err := u.db.SelectInt("select count(*) from Users where username = ? or email = ?", username, email)
	if err != nil {
		return false
	}
	if obj > 0 {
		return true
	}
	return false
}

func NewUserRepository(db *gorp.DbMap) UserRepository {
	u := &user{db}
	return u
}
