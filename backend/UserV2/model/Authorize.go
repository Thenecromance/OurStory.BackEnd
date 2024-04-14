package model

import (
	"github.com/Thenecromance/OurStories/backend/UserV2/data"
	"github.com/Thenecromance/OurStories/base/SQL"
	"github.com/Thenecromance/OurStories/base/hash"
	"github.com/Thenecromance/OurStories/base/log"
)

type Authorization struct {
	tokenCache store
}

func (auth *Authorization) ValidByUserName(username, password string) (usr *data.UserInDb) {
	log.Debug("ValidByUserName: ", username, password)
	usr = &data.UserInDb{}
	err := SQL.Default().SelectOne(usr, "SELECT * FROM user WHERE username=? or email=?", username, username)
	if err != nil {
		log.Error(err)
		return nil
	}

	if hash.Hash(password) != usr.Password {
		return nil
	}
	return
}

func NewAuth() *Authorization {
	return &Authorization{}
}
