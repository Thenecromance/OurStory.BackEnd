package repository

import (
	"fmt"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type HasGetUser interface {
	GetUser(id int64) (*models.User, error)
}
type HasGetUserByUsername interface {
	GetUserByUsername(username string) (*models.User, error)
}
type HasGetUserIdByName interface {
	GetUserIdByName(username string) (int64, error)
}
type HasInsertUser interface {
	InsertUser(user *models.User) error
}
type HasHasUser interface {
	HasUser(username string) bool
}
type HasHasId interface {
	HasId(id int64) bool
}
type HasHasUserAndEmail interface {
	HasUserAndEmail(username, email string) bool
}

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
	HasUser(username string) bool

	HasId(id int64) bool

	// HasUserAndEmail Check if the user and email is exist
	// if the user or email is exist, return true
	// other wise return false
	HasUserAndEmail(username, email string) bool

	// UpdateLastLogin update the last login time of the user
	// deprecated: this method has been implemented by using other method, so this is deprecated
	UpdateLastLogin(id int64, unix int64) error
}

type UserCache interface {
}

const (
	getUserIdByName = "select user_id from Users where username = ?"
	getUserData     = "select * from Users where user_id = ?"
	getAllUser      = "select * from Users"
	getUserByName   = "select * from Users where username = ?"
)

type user struct {
	db *gorp.DbMap

	cache UserCache
}

func (u *user) BindTable() error {
	u.db.AddTableWithName(models.User{}, "Users").SetKeys(true, "UserId")
	u.db.AddTableWithName(models.LoginLogs{}, "LoginLogs").SetKeys(false, "UserId")
	return nil
}

func (u *user) GetUserIdByName(username string) (int64, error) {

	if cache, ok := u.cache.(HasGetUserIdByName); ok {
		id, err := cache.GetUserIdByName(username)
		if id != 0 && err == nil {
			return id, nil
		}
	}

	return u.dbGetUserIdByName(username)
}

func (u *user) dbGetUserIdByName(username string) (int64, error) {
	selectInt, err := u.db.SelectInt(getUserIdByName, username)
	if err != nil {
		return -1, err
	}

	return selectInt, nil
}

func (u *user) GetUser(id int64) (*models.User, error) {
	//models.User
	if cache, ok := u.cache.(HasGetUser); ok {
		usr, err := cache.GetUser(id)
		if usr != nil && err == nil {
			return usr, nil
		}
	}

	return u.dbGetUser(id)
}

func (u *user) dbGetUser(id int64) (*models.User, error) {
	usr := &models.User{}
	err := u.db.SelectOne(usr, getUserData, id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

// GetUsers do not use this by client this will return all data in the table
func (u *user) GetUsers() ([]models.User, error) {
	var users []models.User
	_, err := u.db.Select(&users, getAllUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *user) GetUserByUsername(username string) (*models.User, error) {
	if cache, ok := u.cache.(HasGetUserByUsername); ok {
		usr, err := cache.GetUserByUsername(username)
		if usr != nil && err == nil {
			return usr, nil
		}
	}

	return u.dbGetUserByUsername(username)
}

func (u *user) dbGetUserByUsername(username string) (*models.User, error) {

	obj, err := u.db.Select(models.User{}, getUserByName, username)
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
	if err := u.dbInsertUser(user); err != nil {
		return err
	}

	if cache, ok := u.cache.(HasInsertUser); ok {
		return cache.InsertUser(user)
	}

	return nil
}

func (u *user) dbInsertUser(user *models.User) error {
	err := u.db.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

const (
	checkUserExistsScript     = "select count(*) from Users where username = ?"
	checkUserExistsByIdScript = "select count(*) from Users where user_id = ?"
	checkUserNameAndEmail     = "select count(*) from Users where username = ? and email = ?"
)

func (u *user) HasUser(username string) bool {
	count, err := u.db.SelectInt(checkUserExistsScript, username)
	if err != nil {
		return false
	}

	return count > 0
}

func (u *user) HasId(id int64) bool {
	obj, err := u.db.SelectInt(checkUserExistsByIdScript, id)
	if err != nil {
		return false
	}
	return obj > 0
}

func (u *user) HasUserAndEmail(username, email string) bool {
	obj, err := u.db.SelectInt(checkUserNameAndEmail, username, email)
	if err != nil {
		return false
	}
	return obj > 0
}

func NewUserRepository(db *gorp.DbMap) UserRepository {
	u := &user{
		db:    db,
		cache: newUserRedis(),
	}
	return u
}
