package repository

import (
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/cache/redisCache"
	"github.com/goccy/go-json"
	"strconv"
	"time"
)

var (
	cacheExpireTime = 3600 * time.Second
)

const (
	prefixIdToObject = "user.id"
	prefixNameToId   = "user.name"
	prefixEmailToId  = "user.email"
)

type userRedis struct {
	cli Interface.ICache
}

func (u userRedis) BindTable() error {
	//TODO implement me
	panic("do not use this method with cache, use db instead")
}

func (u userRedis) GetUser(id int64) (*models.User, error) {
	/*//TODO implement me
	panic("implement me")*/
	u.cli.Prefix(prefixIdToObject)
	sId := strconv.FormatInt(id, 10)
	obj, err := u.cli.Get(sId)
	if err != nil {
		return nil, err
	}

	usr := &models.User{}
	err = json.Unmarshal([]byte(obj.(string)), usr)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// GetUsers
// deprecated: cache is not suitable for this method
func (u userRedis) GetUsers() ([]models.User, error) {
	//TODO implement me
	panic("do not use this method with cache, use db instead")
}

func (u userRedis) GetUserByUsername(username string) (*models.User, error) {
	id, err := u.GetUserIdByName(username)
	if err != nil {
		return nil, err
	}
	return u.GetUser(id)
}

func (u userRedis) GetUserIdByName(username string) (int64, error) {
	u.cli.Prefix(prefixNameToId)
	obj, err := u.cli.Get(username)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseInt(obj.(string), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// addIdToUser add user id to cache create a key with prefixIdToObject and value is user object
// key: user.id.{id} value: user object (json)
func (u userRedis) addIdToUser(user *models.User) error {
	u.cli.Prefix(prefixIdToObject)
	sId := strconv.FormatInt(user.UserId, 10)
	buf, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return u.cli.Set(sId, string(buf), cacheExpireTime)
}

// addNameToId add username to cache create a key with prefixNameToId and value is user id
// key: user.name.{username} value: user id
func (u userRedis) addNameToId(user *models.User) error {
	u.cli.Prefix(prefixNameToId)
	return u.cli.Set(user.UserName, strconv.FormatInt(user.UserId, 10), cacheExpireTime)
}

func (u userRedis) addEmailToId(user *models.User) error {
	u.cli.Prefix(prefixEmailToId)
	return u.cli.Set(user.Email, strconv.FormatInt(user.UserId, 10), cacheExpireTime)
}

func (u userRedis) deleteAll(user *models.User) error {
	u.cli.Prefix(prefixIdToObject)
	sId := strconv.FormatInt(user.UserId, 10)
	err := u.cli.Delete(sId)
	if err != nil {
		return err
	}
	u.cli.Prefix(prefixNameToId)
	err = u.cli.Delete(user.UserName)
	if err != nil {
		return err
	}
	u.cli.Prefix(prefixEmailToId)
	return u.cli.Delete(user.Email)

}

func (u userRedis) InsertUser(user *models.User) error {
	err := u.addIdToUser(user)
	if err != nil {
		return err
	}
	err = u.addNameToId(user)
	if err != nil {
		// if error occurred, delete the previous key
		u.cli.Prefix(prefixIdToObject)
		u.cli.Delete(strconv.FormatInt(user.UserId, 10))
		return err
	}
	err = u.addEmailToId(user)
	if err != nil {
		// if error occurred, delete the previous key
		u.cli.Prefix(prefixIdToObject)
		u.cli.Delete(strconv.FormatInt(user.UserId, 10))

		u.cli.Prefix(prefixNameToId)
		u.cli.Delete(user.UserName)
		return err
	}
	return nil
}

func (u userRedis) HasUser(username string) bool {
	id, err := u.GetUserIdByName(username)
	if err != nil {
		return false
	}
	return id > 0
}

func (u userRedis) HasId(id int64) bool {
	usr, err := u.GetUser(id)
	if err != nil {
		return false
	}
	return usr != nil
}

func (u userRedis) hasEmail(email string) bool {
	u.cli.Prefix(prefixEmailToId)
	obj, err := u.cli.Get(email)
	if err != nil {
		return false
	}
	id, err := strconv.ParseInt(obj.(string), 10, 64)
	if err != nil {
		return false
	}
	return id > 0
}

func (u userRedis) HasUserAndEmail(username, email string) bool {
	return u.HasUser(username) || u.hasEmail(email)
}

func (u userRedis) UpdateLastLogin(id int64, unix int64) error {
	panic("do not use this method with cache, use db instead")
}

func newUserRedis() UserRepository {
	return &userRedis{
		cli: redisCache.NewCache(),
	}
}
