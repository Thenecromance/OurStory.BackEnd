package User

import (
	"errors"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/hash"
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
	"time"
)

// when user login or register, or something else need user data, it will be looked like this
type Info struct {
	ID          int    `json:"id" db:"id"`
	UserName    string `json:"username" db:"username, size:20" form:"username" binding:"required"`
	Password    string `json:"password" db:"password, size:64" form:"password" binding:"required"`
	Email       string `json:"email" db:"email, size:64" form:"email"`
	Gender      string `json:"gender" db:"gender"`
	CreatedAt   int    `json:"created_at" db:"created_at"`       // create time stamp
	LastLoginAt int    `json:"last_login_at" db:"last_login_at"` // last login time stamp
	Mate        int    `json:"mate" db:"mate"`                   // after setting the mate, it will be recorded here
}

type Model struct {
	db *gorp.DbMap
	// the place where stored logged in user
	cache pool
}

func (m *Model) initdb() error {
	if m.db != nil {
		return nil
	}
	m.db = SQL.Default()
	tm := m.db.AddTableWithName(Info{}, "users")
	tm.SetKeys(true, "ID")
	tm.ColMap("username").SetUnique(true).SetNotNull(true)
	tm.ColMap("email").SetUnique(true).SetNotNull(true)
	tm.ColMap("password").SetNotNull(true)
	tm.ColMap("gender")

	err := m.db.CreateTablesIfNotExists()
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	m.db.AddTableWithName(loginInfo{}, "loginInfo").SetKeys(true, "ID")
	err = m.db.CreateTablesIfNotExists()
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	return nil
}

func (m *Model) getUserFromDatabase(user *Info) Info {
	logger.Get().Debugf("start to get user %s from database", user.UserName)
	err := m.initdb()
	if err != nil {

		return Info{}
	}

	var temp Info
	err = m.db.SelectOne(&temp, "SELECT * FROM `users` WHERE username = ? or email = ?", user.UserName, user.Email)
	if err != nil {
		return Info{}
	}
	return temp
}

func (m *Model) userInDatabase(user *Info) bool {
	return m.getUserFromDatabase(user).ID != 0
}

// -----------------------------------------------------------------

func (m *Model) recordUserLoginTime(userId int) {
	logged := loginInfo{
		UserId:    userId,
		TimeStamp: time.Now().Unix(),
	}
	m.db.Insert(&logged)
}

func (m *Model) register(user Info) error {
	//pre check user already exists
	{
		logger.Get().Infof("%s start to register", user.UserName)
		// check if user already loaded in to the cache
		if m.cache.getByName(user.UserName).ID != 0 {
			logger.Get().Infof("%s already exists", user.UserName)
			return errors.New("user already exists")
		}
		// check if user already in the database
		if m.userInDatabase(&user) {
			logger.Get().Infof("%s already exists", user.UserName)
			return errors.New("user already exists")
		}
	}

	user.CreatedAt = int(time.Now().Unix())
	user.LastLoginAt = user.CreatedAt
	user.Password = hash.Hash(user.Password)
	err := m.db.Insert(&user)
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	// update user info to cache
	user = m.getUserFromDatabase(&user)
	//push to cache
	m.cache.add(user)

	m.recordUserLoginTime(user.ID)

	return nil
}

func (m *Model) login(requestUser *Info) error {
	userInDatabase := m.cache.getByName(requestUser.UserName)
	if userInDatabase.ID == 0 {
		userInDatabase = m.getUserFromDatabase(requestUser)
	}

	if userInDatabase.ID == 0 {
		logger.Get().Infof("requestUser not exists")
		return errors.New("requestUser not exists")
	}
	if userInDatabase.UserName != requestUser.UserName {
		logger.Get().Info("username not match")
		return errors.New("username not match")
	}
	if userInDatabase.Password != hash.Hash(requestUser.Password) {
		logger.Get().Info("password not match")
		return errors.New("password not match")
	}

	requestUser.ID = userInDatabase.ID

	m.cache.add(*requestUser)
	logger.Get().Infof("%s login success", requestUser.UserName)

	m.recordUserLoginTime(requestUser.ID)
	return nil
}

//-----------------------------------------------------------------
