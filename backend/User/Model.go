package User

import (
	"errors"
	"github.com/Thenecromance/OurStories/base/hash"
	"github.com/Thenecromance/OurStories/base/logger"
	"time"
)

type Model struct {
	//db *gorp.DbMap
	// the place where stored logged in user
	cache pool
}

func (m *Model) getUserFromDatabase(user *Info) Info {
	/*logger.Get().Debugf("start to get user %s from database", user.UserName)
	err := m.initdb()
	if err != nil {

		return Info{}
	}

	var temp Info
	err = m.db.SelectOne(&temp, "SELECT * FROM `users` WHERE username = ? or email = ?", user.UserName, user.Email)
	if err != nil {
		return Info{}
	}
	return temp*/
	var temp Info
	temp.GetUserFromDatabase(user)
	return temp
}

func (m *Model) userInDatabase(user *Info) bool {
	return m.getUserFromDatabase(user).ID != 0
}

// -----------------------------------------------------------------

func (m *Model) recordUserLoginTime(userId int) {
	loggedinfo := loginInfo{
		UserId:    userId,
		TimeStamp: time.Now().Unix(),
	}
	loggedinfo.SelfInsert()
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

	{
		user.CreatedAt = int(time.Now().Unix())  // record the create time
		user.LastLoginAt = user.CreatedAt        // record the last login time
		user.Password = hash.Hash(user.Password) // hash the password
	}

	err := user.SelfInsert()
	if err != nil {
		logger.Get().Error(err)
		return err
	}

	// update user self info
	err = user.SelfGet()
	if err != nil {
		logger.Get().Error(err)
		return err
	}
	//push to cache
	m.cache.add(user)

	m.recordUserLoginTime(user.ID)
	return nil
}

func (m *Model) login(requestUser *Info) error {

	// check user is in the cache or not
	userInDatabase := m.cache.getByName(requestUser.UserName)
	if userInDatabase.ID == 0 {
		// if user does not exist in the cache, so I need to get it from the database
		userInDatabase = m.getUserFromDatabase(requestUser)
	}

	//user check sections
	// if user's id is 0 means user not exists
	if userInDatabase.ID == 0 {
		logger.Get().Infof("requestUser not exists")
		return errors.New("requestUser not exists")
	}
	//compare the username and pass word
	if userInDatabase.UserName != requestUser.UserName {
		logger.Get().Info("username not match")
		return errors.New("username not match")
	}
	if userInDatabase.Password != hash.Hash(requestUser.Password) {
		logger.Get().Info("password not match")
		return errors.New("password not match")
	}

	{
		requestUser.ID = userInDatabase.ID
		requestUser.UserName = userInDatabase.UserName
		requestUser.Password = userInDatabase.Password
		requestUser.Email = userInDatabase.Email
		requestUser.Gender = userInDatabase.Gender
		requestUser.CreatedAt = userInDatabase.CreatedAt
		requestUser.LastLoginAt = userInDatabase.LastLoginAt
		requestUser.Mate = userInDatabase.Mate
	}

	//just copy
	m.cache.add(userInDatabase)

	logger.Get().Infof("%s login success", requestUser.UserName)

	m.recordUserLoginTime(requestUser.ID)
	return nil
}

//-----------------------------------------------------------------
