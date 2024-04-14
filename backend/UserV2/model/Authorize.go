package model

import (
	"github.com/Thenecromance/OurStories/backend/UserV2/data"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/hash"
	"github.com/Thenecromance/OurStories/base/log"
	"github.com/Thenecromance/OurStories/base/lru"
	"time"
)

type Authorization struct {
	usrCache   *lru.Cache
	tokenCache store
}

func (auth *Authorization) ValidByUserName(username, password string) (usr *data.UserInDb) {

	usr = &data.UserInDb{}
	ptr, ok := auth.usrCache.Get(username)
	if !ok {
		err := SQL.Default().SelectOne(usr, "SELECT * FROM user WHERE username=? or email=?", username, username)
		if err != nil {
			log.Error(err)
			return nil
		}
		auth.usrCache.Add(username, usr, time.Now().Add(15*time.Minute))
	} else {
		usr = ptr.(*data.UserInDb)
	}

	if hash.Hash(password) != usr.Password {
		return nil
	}
	return
}

func NewAuth() *Authorization {

	return &Authorization{
		usrCache: lru.New(100),
	}
}
